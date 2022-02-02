package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/poroping/go-cisco-umbrella-sdk/models"
)

func (c *Client) UpdateApiKey(serviceProviderId, customerId int, apiKeyId string) (*models.ApiKeyResp, error) {
	apiKeyPath := fmt.Sprintf("/v1/serviceproviders/%d/customers/%d/apikeys", serviceProviderId, customerId)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s%s/%s", c.HostURL, apiKeyPath, apiKeyId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	res := models.ApiKeyResp{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateApiKey(serviceProviderId, customerId int) (*models.ApiKeyResp, error) {
	apiKeyPath := fmt.Sprintf("/v1/serviceproviders/%d/customers/%d/apikeys", serviceProviderId, customerId)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.HostURL, apiKeyPath), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	res := models.ApiKeyResp{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadApiKey(serviceProviderId, customerId int, apiKeyId string) (*models.ApiKeyResp, error) {
	apiKeyPath := fmt.Sprintf("/v1/serviceproviders/%d/customers/%d/apikeys", serviceProviderId, customerId)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s/%s", c.HostURL, apiKeyPath, apiKeyId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	res := models.ApiKeyResp{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteApiKey(serviceProviderId, customerId int, apiKeyId string) error {
	apiKeyPath := fmt.Sprintf("/v1/serviceproviders/%d/customers/%d/apikeys", serviceProviderId, customerId)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s/%s", c.HostURL, apiKeyPath, apiKeyId), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ListApiKey(serviceProviderId, customerId int) (*[]models.ApiKeyResp, error) {
	apiKeyPath := fmt.Sprintf("/v1/serviceproviders/%d/customers/%d/apikeys", serviceProviderId, customerId)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.HostURL, apiKeyPath), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	res := []models.ApiKeyResp{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
