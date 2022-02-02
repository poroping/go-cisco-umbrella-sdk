package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/poroping/go-cisco-umbrella-sdk/models"
)

const contentType = "application/json"
const managementBase = "https://management.api.umbrella.com"

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	apiID      string
	apiKey     string
	userAgent  string
}

func NewClient(apiID, apiKey, userAgent string, insecure bool) (*Client, error) {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}

	retryClient.HTTPClient.Transport = tr
	retryClient.HTTPClient.Timeout = 10 * time.Second

	c := Client{
		HTTPClient: retryClient.StandardClient(),
		HostURL:    managementBase,
		apiID:      apiID,
		apiKey:     apiKey,
		userAgent:  userAgent,
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", contentType)
	req.Header.Set("User-Agent", c.userAgent)
	req.SetBasicAuth(c.apiID, c.apiKey)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] Status code: %d", res.StatusCode)
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		errres := models.ErrorResp{}
		err = json.Unmarshal(body, &errres)
		if err != nil {
			return nil, err
		}
		e := ""
		if errres.Error != nil {
			e = *errres.Error
		}
		msg := ""
		if errres.Message != nil {
			msg = *errres.Message
		}
		return nil, fmt.Errorf("error: %s | message: %s", e, msg)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
