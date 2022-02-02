package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/poroping/go-cisco-umbrella-sdk/models"
)

const serviceProviderPath = "/v1/serviceproviders/"

func (c *Client) ReadServiceProvider(serviceProviderId int) (*models.ServiceProviderResp, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s%d", c.HostURL, serviceProviderPath, serviceProviderId), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	res := models.ServiceProviderResp{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
