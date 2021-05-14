package statusflare

import "fmt"

type Integration struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

func (c *Client) AllIntegrations() ([]*Integration, error) {
	url := fmt.Sprintf("/integrations/%s", c.accountID)
	resp, err := c.makeAPICall("GET", url, nil)
	if err != nil {
		return []*Integration{}, err
	}

	integrations := []*Integration{}
	err = unmarshallResp(resp, &integrations)

	return integrations, err
}
