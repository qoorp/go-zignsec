package zignsec

// This is an interface to http://docs.zignsec.com/api/web-based/

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	// APIHostBase is the production endpoint
	APIHostBase = "https://api.zignsec.com/v2/bankidse"
	// APIHostBaseTest is the test endpoint
	APIHostBaseTest = "https://test.zignsec.com/v2/bankidse"
)

// Client is a Zignsec web-based client.
type Client struct {
	APIHostBase string
	APIKey      string
}

// New create a new Client
func New(APIHostBase string, APIKey string) *Client {
	c := new(Client)
	c.APIHostBase = APIHostBase
	c.APIKey = APIKey
	return c
}

// Initiate a login or sign request
func (c *Client) Initiate(method string, config ZSInitConfig) (*ZSInitRespBody, error) {
	var url string
	url = c.APIHostBase + "/" + method
	configB, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(configB))
	if err != nil {
		return nil, err
	}
	setHeaders(req, c.APIKey, "application/json")
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
	var response ZSInitRespBody
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// Verify a login or signature
func (c *Client) Verify(orderRef string) (*ZSVerifyRespBody, error) {
	url := c.APIHostBase + "/collect?orderref=" + orderRef
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req, c.APIKey, "application/x-www-form-urlencoded")
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
	var response ZSVerifyRespBody
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// Cancel a login or signature
func (c *Client) Cancel(orderRef string) (*ZSCancelResponse, error) {
	url := c.APIHostBase + "/cancel"
	cancelBody := ZSCancelBody{OrderRef: orderRef}
	body, err := json.Marshal(cancelBody)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	setHeaders(req, c.APIKey, "application/json")
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

func setHeaders(req *http.Request, APIKey string, contentType string) {
	req.Header.Add("Authorization", APIKey)
	req.Header.Add("Content-Type", contentType)
}
