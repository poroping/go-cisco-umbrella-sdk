Stuff

| License Type | Name |
|---|---|
| `term` | term customer | 
| `msla` | msla customer |

| Package Id | Package Name |
|---|---|
| `250` | Umbrella SIG Essentials | 
| `248` | Umbrella DNS Security Advantage |
| `246` | Umbrella DNS Security Essentials |
| `202` | Cisco Umbrella for EDU |
| `171` | Cisco Umbrella for Wireless LAN |
| `107` | Umbrella Insights |
| `101` | Umbrella Platform |
| `99` | Umbrella Professional |

### Usage 

```go
package main

import (
	"fmt"

	"github.com/poroping/go-cisco-umbrella-sdk/client"
	"github.com/poroping/go-cisco-umbrella-sdk/models"
)

const apiID = "asdasdas"
const apiKey = "asdas"
const serviceProviderId = 123456
const customerId = 12345656

func main() {
	c, err := client.NewClient(apiID, apiKey, "go-cisco-umbrella-sdk", false)
	if err != nil {
		log.Fatalf("[FATAL] %s", err)
	}
	resp, err := c.ReadServiceProvider(serviceProviderId)
	if err != nil {
		log.Fatalf("[FATAL] %s", err)
	}
	fmt.Println(resp)

	listCust, err := c.ListCustomer(serviceProviderId)
	if err != nil {
		log.Fatalf("[FATAL] %s", err)
	}
	fmt.Println(listCust)

	readCust, err := c.ReadCustomer(serviceProviderId, customerId)
	if err != nil {
		log.Fatalf("[FATAL] %s", err)
	}
	fmt.Println(readCust)

	t := true
	licenseType := "msla"

	newCust := models.CustomerParams{}
	newCust.IsTrial = &t
	newCust.CustomerName = "Pen Island Inc."
	newCust.Seats = 1
	newCust.PackageID = 246
	newCust.LicenseType = &licenseType
	newCust.StreetAddress = "420 Noice Street"
	newCust.City = "Reykjavik"
	newCust.State = "Reykjavik"
	newCust.CountryCode = "IS"
	newCust.AdminEmails = []string{"email@emails.com"}

	newCustomer, err := c.CreateCustomer(serviceProviderId, newCust)
	if err != nil {
		log.Fatalf("[FATAL] %s", err)
	}
	fmt.Println(newCustomer.CustomerID)

	newCust.Seats = 2
	updateCust, err := c.UpdateCustomer(serviceProviderId, newCustomer.CustomerID, newCust)
	if err != nil {
		log.Fatalf("[FATAL] %s", err)
	}
	fmt.Println(updateCust.Seats)

	apiKey, err := c.CreateApiKey(serviceProviderId, newCustomer.CustomerID)
	if err != nil {
		log.Fatalf("[FATAL] %s", err)
	}
	fmt.Println(apiKey.Key)

	apiKeyUpd, err := c.UpdateApiKey(serviceProviderId, newCustomer.CustomerID, apiKey.KeyID)
	if err != nil {
		log.Fatalf("[FATAL] %s", err)
	}
	fmt.Println(apiKeyUpd.Key)

	err = c.DeleteCustomer(serviceProviderId, newCustomer.CustomerID)
	if err != nil {
		log.Fatalf("[FATAL] %s", err)
	}
}

```