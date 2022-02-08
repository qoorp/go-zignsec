// Server to server (S2S) API from Zignsec
// https://docs.zignsec.com/api/v2/s2s/

package zignsec

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	// S2SURL is the production endpoint
	S2SURL = "https://api.zignsec.com/v2/BankIDSE"
	// S2SURLTest is the test endpoint
	S2SURLTest = "https://test.zignsec.com/v2/BankIDSE"
)

// S2SClient is a Zignsec server to server client
type S2SClient struct {
	baseURL string
	key     string
}

// NewS2SClient create a new server to server client.
// Key is the Zignsec Authorization.
func NewS2SClient(baseURL string, key string) *S2SClient {
	result := S2SClient{baseURL: baseURL, key: key}
	return &result
}

// Collect the state of an authenticate or sign request.
// Order is from the Init() response.
// The guidelines recommend polling for results every 2 seconds.
// When the response Status is COMPLETE, then further use of Collect() will receive an error.
func (c *S2SClient) Collect(orderRef string) (*ZSCollectResponse, error) {
	var result ZSCollectResponse

	url := c.baseURL + "/Collect?orderRef=" + orderRef
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
func (c *S2SClient) Init(method string, config ZSInitConfig) (*ZSInitResponse, error) {
	var result ZSInitResponse

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

// Cancel a login or signature
func (c *S2SClient) Cancel(orderRef string) (*ZSCancelResponse, error) {
	url := c.baseURL + "/cancel"
	cancelBody := ZSCancelBody{OrderRef: orderRef}
	body, err := json.Marshal(cancelBody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
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
	var response ZSCancelResponse
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
