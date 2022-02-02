package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/poroping/go-cisco-umbrella-sdk/models"
)

func (c *Client) UpdateCustomer(serviceProviderId, customerId int, params models.CustomerParams) (*models.CustomerResp, error) {
	customerPath := fmt.Sprintf("/v1/serviceproviders/%d/customers", serviceProviderId)

	// null fields not updateable
	params.IsTrial = nil
	params.LicenseType = nil

	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s%s/%d", c.HostURL, customerPath, customerId), strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Request Body: %s", string(body))

	res := models.CustomerResp{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateCustomer(serviceProviderId int, params models.CustomerParams) (*models.CustomerResp, error) {
	customerPath := fmt.Sprintf("/v1/serviceproviders/%d/customers", serviceProviderId)

	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.HostURL, customerPath), strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Request Body: %s", string(body))

	res := models.CustomerResp{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReadCustomer(serviceProviderId, customerId int) (*models.CustomerResp, error) {
	customerPath := fmt.Sprintf("/v1/serviceproviders/%d/customers", serviceProviderId)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s/%d", c.HostURL, customerPath, customerId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	res := models.CustomerResp{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteCustomer(serviceProviderId, customerId int) error {
	customerPath := fmt.Sprintf("/v1/serviceproviders/%d/customers", serviceProviderId)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s/%d", c.HostURL, customerPath, customerId), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ListCustomer(serviceProviderId int) (*[]models.CustomerResp, error) {
	customerPath := fmt.Sprintf("/v1/serviceproviders/%d/customers", serviceProviderId)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.HostURL, customerPath), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	res := []models.CustomerResp{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
