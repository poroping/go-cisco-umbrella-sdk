package models

type ServiceProviderResp struct {
	PackageName   string `json:"packageName"`
	SeatsTotal    int    `json:"seatsTotal"`
	SeatsUsed     int    `json:"seatsUsed"`
	CustomerCount int    `json:"customerCount"`
	Status        string `json:"status"`
	RebillAt      string `json:"rebillAt"`
	ExpiresAt     string `json:"expiresAt"`
}
