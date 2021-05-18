package statusflare

import (
	"encoding/json"
	"fmt"
)

type StatusPage struct {
	Id                 string           `json:"id"`
	Name               string           `json:"name"`
	Monitors           []string         `json:"monitors"`
	CustomDomain       string           `json:"custom_domain,omitempty"` // not sure if it's typo in API, but I'm using Scheme here
	CustomDomainPath   string           `json:"custom_domain_path,omitempty"`
	HideMonitorDetails bool             `json:"hide_monitor_details"`
	HideStatusflare    bool             `json:"hide_statusflare"`
	Config             StatusPageConfig `json:"config,omitempty"`
}

type StatusPageConfig struct {
	Title                      string `json:"title,omitempty"`
	HistogramDays              int    `json:"histogram_days,omitempty"`
	LogoUrl                    string `json:"logo_url,omitempty"`
	FaviconUrl                 string `json:"favicon_url,omitempty"`
	AllMonitorsOperational     string `json:"all_monitors_operational,omitempty"`
	NotAllMonitorsOperational  string `json:"not_all_monitors_operational,omitempty"`
	MonitorOperationalLabel    string `json:"monitor_operational_label,omitempty"`
	MonitorNotOperationalLabel string `json:"monitor_not_operational_label,omitempty"`
	MonitorNoDataLabel         string `json:"monitor_no_data_label,omitempty"`
	HistogramNoData            string `json:"histogram_no_data,omitempty"`
	HistogramNoIncidents       string `json:"histogram_no_incidents,omitempty"`
	HistogramSomeIncidents     string `json:"histogram_some_incidents,omitempty"`
}

// create new status page
//
// When create process is successful, the
// function populate given 's' by new values like ID
// etc.
//
// Be aware you're usig API token wih read & write
// permissions.
func (c *Client) CreateStatusPage(s *StatusPage) error {
	var err error

	body, err := json.Marshal(s)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("/status-pages/%s", c.accountId)
	resp, err := c.makeAPICall("POST", url, body)
	if err != nil {
		return err
	}

	return unmarshallResp(resp, s)
}

// returns list of all status pages
// the pagination is currently questionable. Let's
// assume function gives you all status pages
func (c *Client) AllStatusPages() ([]*StatusPage, error) {
	url := fmt.Sprintf("/status-pages/%s", c.accountId)
	resp, err := c.makeAPICall("GET", url, nil)
	if err != nil {
		return []*StatusPage{}, err
	}

	statusPages := []*StatusPage{}
	err = unmarshallResp(resp, &statusPages)

	return statusPages, err
}

// get status page for given ID.
// If there is no status page for ID, the error is returned.
func (c *Client) GetStatusPage(id string) (*StatusPage, error) {
	var statusPage StatusPage

	url := fmt.Sprintf("/status-pages/%s/%s", c.accountId, id)
	resp, err := c.makeAPICall("GET", url, nil)
	if err != nil {
		return &statusPage, err
	}

	err = unmarshallResp(resp, &statusPage)
	return &statusPage, err
}

// update the existing status page in statusflare.
//
// This function require presense of values in 's'
// like accountID or ID. You can use status page returned
// by GetStatusPAge or AllStatusPages
//
// Also 's' will be populated  by values the update returns
func (c *Client) SaveStatusPage(s *StatusPage) error {
	var err error

	body, err := json.Marshal(s)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("/status-pages/%s/%s", c.accountId, s.Id)
	resp, err := c.makeAPICall("PUT", url, body)
	if err != nil {
		return err
	}

	return unmarshallResp(resp, s)
}

// function delete the status page in statusflare.
func (c *Client) DeleteStatusPage(id string) error {
	url := fmt.Sprintf("/status-pages/%s/%s", c.accountId, id)
	_, err := c.makeAPICall("DELETE", url, nil)
	if err != nil {
		return err
	}

	return nil
}
