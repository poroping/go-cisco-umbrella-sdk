package models

type CustomerParams struct {
	CustomerName       string   `json:"customerName,omitempty"`
	LicenseType        *string  `json:"licenseType,omitempty"`
	IsTrial            *bool    `json:"isTrial,omitempty"`
	Seats              int      `json:"seats,omitempty"`
	StreetAddress      string   `json:"streetAddress,omitempty"`
	StreetAddress2     string   `json:"streetAddress2,omitempty"`
	City               string   `json:"city,omitempty"`
	State              string   `json:"state,omitempty"`
	CountryCode        string   `json:"countryCode,omitempty"`
	ZipCode            string   `json:"zipCode,omitempty"`
	PackageID          int      `json:"packageId,omitempty"`
	DealID             string   `json:"dealId,omitempty"`
	AdminEmails        []string `json:"adminEmails,omitempty"`
	CcwDealOwnerEmails []string `json:"ccwDealOwnerEmails,omitempty"`
}

type CustomerResp struct {
	CustomerName       string   `json:"customerName"`
	LicenseType        string   `json:"licenseType"`
	IsTrial            bool     `json:"isTrial"`
	Seats              int      `json:"seats"`
	StreetAddress      string   `json:"streetAddress"`
	StreetAddress2     string   `json:"streetAddress2"`
	City               string   `json:"city"`
	State              string   `json:"state"`
	CountryCode        string   `json:"countryCode"`
	ZipCode            string   `json:"zipCode"`
	PackageID          int      `json:"packageId"`
	DealID             string   `json:"dealId"`
	AdminEmails        []string `json:"adminEmails"`
	CcwDealOwnerEmails []string `json:"ccwDealOwnerEmails"`
	CreatedAt          string   `json:"createdAt"`
	ModifiedAt         string   `json:"modifiedAt"`
	CustomerID         int      `json:"customerId"`
}
