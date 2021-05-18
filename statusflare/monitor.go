package statusflare

import (
	"encoding/json"
	"fmt"
)

type Monitor struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	URL                string   `json:"url"`
	Scheme             string   `json:"schema"` // not sure if it's typo in API, but I'm using Scheme here
	Method             string   `json:"method"`
	ExpectStatus       int      `json:"expect_status"`
	NotifyAfter        int      `json:"notify_after"`
	Worker             string   `json:"worker"`
	Integrations       []string `json:"integrations"`
	FollowRedirects    bool     `json:"follow_redirects"`
	InsecureSkipVerify bool     `json:"insecure_skip_verify"`
	Timeout            int      `json:"timeout"`
	Interval           int      `json:"interval"`
}

// create new monitor
//
// When create process is successful, the
// function populate given 'm' by new values like ID
// etc.
//
// Be aware you're usig API token wih read & write
// permissions.
func (c *Client) CreateMonitor(m *Monitor) error {
	var err error

	body, err := json.Marshal(m)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("/uptime/%s", c.accountId)
	resp, err := c.makeAPICall("POST", url, body)
	if err != nil {
		return err
	}

	return unmarshallResp(resp, m)
}

// returns list of all monitors
// the pagination is currently questionable. Let's
// assume function gives you all monitors
func (c *Client) AllMonitors() ([]*Monitor, error) {
	url := fmt.Sprintf("/uptime/%s", c.accountId)
	resp, err := c.makeAPICall("GET", url, nil)
	if err != nil {
		return []*Monitor{}, err
	}

	monitors := []*Monitor{}
	err = unmarshallResp(resp, &monitors)

	return monitors, err
}

// get monitor for given ID.
// If there is no monitor for ID, the error is returned.
func (c *Client) GetMonitor(id string) (*Monitor, error) {
	var monitor Monitor

	url := fmt.Sprintf("/uptime/%s/%s", c.accountId, id)
	resp, err := c.makeAPICall("GET", url, nil)
	if err != nil {
		return &monitor, err
	}

	err = unmarshallResp(resp, &monitor)
	return &monitor, err
}

// update the existing monitor in statusflare.
//
// This function require presense of values in 'm'
// like accountID or ID. You can use monitor returned
// by GetMonitor or AllMonitors
//
// Also 'm' will be populated  by values the update returns
func (c *Client) SaveMonitor(m *Monitor) error {
	var err error

	body, err := json.Marshal(m)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("/uptime/%s/%s", c.accountId, m.Id)
	resp, err := c.makeAPICall("PUT", url, body)
	if err != nil {
		return err
	}

	return unmarshallResp(resp, m)
}

// function delete the monitor in statusflare.
func (c *Client) DeleteMonitor(id string) error {
	url := fmt.Sprintf("/uptime/%s/%s", c.accountId, id)
	_, err := c.makeAPICall("DELETE", url, nil)
	if err != nil {
		return err
	}

	return nil
}
