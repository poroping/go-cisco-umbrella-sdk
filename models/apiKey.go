package models

type ApiKeyResp struct {
	Name       string `json:"name"`
	KeyID      string `json:"keyId"`
	CustomerID string `json:"customerId"`
	Key        string `json:"key"`
	Secret     string `json:"secret"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}
