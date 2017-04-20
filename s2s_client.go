// Server to server (S2S) API from Zignsec

package zignsec

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// S2SURL is the production endpoint
	S2SURL = "https://api.zignsec.com/v2/BankIDSE"
	// S2SURLTest is the test endpoint
	S2SURLTest = "https://test.zignsec.com/v2/BankIDSE"
	// Timeout3 is how long BankID will wait before cancelling a request
	Timeout3 = time.Duration(3 * time.Minute)
)

// S2SClient is a Zignsec server to server client
type S2SClient struct {
	baseURL string
	key     string
	timeout time.Duration
}

// NewS2SClient create a new client.
// Key is the Zignsec Authorization.
// Timeout is only used by Authenticate() and Sign().
func NewS2SClient(baseURL string, key string, timeout time.Duration) *S2SClient {
	result := S2SClient{baseURL: baseURL, key: key, timeout: timeout}
	return &result
}

// Authenticate request that does both Init() and Collect().
func (c *S2SClient) Authenticate(nr string) (*CollectResponse, error) {
	config := ZSInitConfig{Personalnumber: nr}
	init, err := c.Init("Authenticate", config)
	if err != nil {
		return nil, err
	}
	return c.loop(init)
}

// Collect the state of an authenticate or sign request.
// Order is from the Init() response.
// This step is included in Authenticate() and Sign().
// The guidelines recommend polling for results every 2 seconds.
// When the response Status is COMPLETE, then further use of Collect() will receive an error.
func (c *S2SClient) Collect(order string) (*CollectResponse, error) {
	var result CollectResponse

	url := c.baseURL + "/Collect?orderRef=" + order
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.key)
	var httpClient http.Client
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	err = json.Unmarshal(b, &result)
	return &result, err
}

// Init a authenticate or sign request.
// Use Collect() to get the answer afterwards.
// This step is included in Authenticate() and Sign().
func (c *S2SClient) Init(method string, config ZSInitConfig) (*InitResponse, error) {
	var result InitResponse

	url := c.baseURL + "/" + method
	configJSON, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(configJSON))
	if err != nil {
		return nil, err
	}
	setHeaders(req, c.key, "application/json")
	var httpClient http.Client
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Sign request that does both Init() and Collect().
// It is sign that is signed. Prompt is for user guidance.
// If prompt contains certain bytes (> 128?), Init() will fail with INVALID_PARAMETERS.
func (c *S2SClient) Sign(Personalnumber string, prompt, sign []byte) (*CollectResponse, error) {
	v := base64.StdEncoding.EncodeToString(prompt)
	inv := base64.StdEncoding.EncodeToString(sign)
	config := ZSInitConfig{Personalnumber: Personalnumber, UserVisibleData: v, UserNonVisibleData: inv}
	init, err := c.Init("Sign", config)
	if err != nil {
		return nil, err
	}
	return c.loop(init)
}

//
// Internal functions
//

func (c *S2SClient) loop(init *InitResponse) (*CollectResponse, error) {
	var result *CollectResponse

	if len(init.Errors) != 0 {
		return nil, errors.New(init.Errors[0].Description)
	}
	until := time.Now().Add(c.timeout)
	for time.Now().Before(until) {
		time.Sleep(2 * time.Second)
		result, err := c.Collect(init.Order)
		if err != nil {
			return result, err
		}
		if len(result.Errors) != 0 {
			return result, errors.New(result.Errors[0].Description)
		}
		if result.Status == "COMPLETE" {
			return result, nil
		}
	}
	return result, errors.New("timeout")
}
