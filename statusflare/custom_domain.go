package statusflare

import (
	"encoding/json"
	"fmt"
)

type CustomDomain struct {
	Id     string `json:"id"`
	Domain string `json:"domain"`
	Type   string `json:"type"`
	Custom bool   `json:"custom"`
	Status string `json:"status"`
}

// create new custom domain
//
// When create process is successful, the
// function populate given 'd' by new values like ID
// etc.
//
// Be aware you're usig API token wih read & write
// permissions.
func (c *Client) CreateCustomDomain(d *CustomDomain) error {
	var err error

	body, err := json.Marshal(d)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("/account/custom-domains/%s", c.accountId)
	resp, err := c.makeAPICall("POST", url, body)
	if err != nil {
		return err
	}

	return unmarshallResp(resp, d)
}

// returns list of all custom domains
// the pagination is currently questionable. Let's
// assume function gives you all custom domains
func (c *Client) AllCustomDomains() ([]*CustomDomain, error) {
	url := fmt.Sprintf("/account/custom-domains/%s", c.accountId)
	resp, err := c.makeAPICall("GET", url, nil)
	if err != nil {
		return []*CustomDomain{}, err
	}

	customDomains := []*CustomDomain{}
	err = unmarshallResp(resp, &customDomains)

	return customDomains, err
}

// get custom domain for given ID.
// If there is no custom domain for ID, the error is returned.
func (c *Client) GetCustomDomain(id string) (*CustomDomain, error) {
	var customDomain CustomDomain

	url := fmt.Sprintf("/account/custom-domains/%s/%s", c.accountId, id)
	resp, err := c.makeAPICall("GET", url, nil)
	if err != nil {
		return &customDomain, err
	}

	err = unmarshallResp(resp, &customDomain)
	return &customDomain, err
}

// function delete the custom domain in statusflare.
func (c *Client) DeleteCustomDomain(id string) error {
	url := fmt.Sprintf("/account/custom-domains/%s/%s", c.accountId, id)
	_, err := c.makeAPICall("DELETE", url, nil)
	if err != nil {
		return err
	}

	return nil
}
