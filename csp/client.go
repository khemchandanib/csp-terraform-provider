package csp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const HostURL string = "http://localhost:19090"

// CspClient -
type CspClient struct {
	HostURL             string
	ServiceDefinitionID string
	HTTPClient          *http.Client
	Token               string
}

// NewCspClient -
func NewCspClient(host, token, serviceDefinitionId *string) (*CspClient, error) {
	c := CspClient{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	// If token not provided, return empty client
	if token == nil {
		return &c, nil
	}

	c.Token = *token

	if serviceDefinitionId == nil {
		return &c, nil
	}

	c.ServiceDefinitionID = *serviceDefinitionId

	return &c, nil
}

func (c *CspClient) doRequest(req *http.Request) ([]byte, error) {
	token := c.Token

	req.Header.Set("Authorization", "Bearer "+token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
