package statusflare

import (
	"encoding/json"
	"fmt"
)

type Integration struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Secret string `json:"secret,omitempty"`
}

// create new integration
//
// When create process is successful, the
// function populate given 'i' by new values like ID
// etc.
//
// Be aware you're usig API token wih read & write
// permissions.
func (c *Client) CreateIntegration(i *Integration) error {
	var err error

	body, err := json.Marshal(i)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("/integrations/%s", c.accountId)
	resp, err := c.makeAPICall("POST", url, body)
	if err != nil {
		return err
	}

	return unmarshallResp(resp, i)
}

// returns list of all integrations
// the pagination is currently questionable. Let's
// assume function gives you all integrations
func (c *Client) AllIntegrations() ([]*Integration, error) {
	url := fmt.Sprintf("/integrations/%s", c.accountId)
	resp, err := c.makeAPICall("GET", url, nil)
	if err != nil {
		return []*Integration{}, err
	}

	integrations := []*Integration{}
	err = unmarshallResp(resp, &integrations)

	return integrations, err
}

// get integration for given ID.
// If there is no integration for ID, the error is returned.
func (c *Client) GetIntegration(id string) (*Integration, error) {
	var integration Integration

	url := fmt.Sprintf("/integrations/%s/%s", c.accountId, id)
	resp, err := c.makeAPICall("GET", url, nil)
	if err != nil {
		return &integration, err
	}

	err = unmarshallResp(resp, &integration)
	return &integration, err
}

// update the existing integration in statusflare.
//
// This function require presense of values in 'i'
// like accountID or ID. You can use integration returned
// by GetIntegration or AllIntegrations
//
// Also 'i' will be populated  by values the update returns
func (c *Client) SaveIntegration(i *Integration) error {
	var err error

	body, err := json.Marshal(i)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("/integrations/%s/%s", c.accountId, i.Id)
	resp, err := c.makeAPICall("PUT", url, body)
	if err != nil {
		return err
	}

	return unmarshallResp(resp, i)
}

// function delete the monitor in statusflare.
func (c *Client) DeleteIntegration(id string) error {
	url := fmt.Sprintf("/integrations/%s/%s", c.accountId, id)
	_, err := c.makeAPICall("DELETE", url, nil)
	if err != nil {
		return err
	}

	return nil
}
