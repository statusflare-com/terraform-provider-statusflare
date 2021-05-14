package statusflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// represent the session to statusflare
type Client struct {
	url       string
	accountID string
	keyID     string
	token     string
	http      *http.Client
}

// create default client where all needed data are
// read from env. variables. The required env. variables
// are:
//   - STATUSFLARE_ACCOUNT_ID
//   - STATUSFLARE_KEY_ID
//   - STATUSFLARE_TOKEN
//
// If none of the env. variables are set, the function returns
// you nil client and error.
func DefaultClient() (*Client, error) {

	if os.Getenv("STATUSFLARE_ACCOUNT_ID") == "" {
		return nil, fmt.Errorf("missing STATUSFLARE_ACCOUNT_ID")
	}

	if os.Getenv("STATUSFLARE_KEY_ID") == "" {
		return nil, fmt.Errorf("missing STATUSFLARE_KEY_ID")
	}

	if os.Getenv("STATUSFLARE_TOKEN") == "" {
		return nil, fmt.Errorf("missing STATUSFLARE_TOKEN")
	}

	client := &Client{
		url:       "https://api.statusflare.com/",
		accountID: os.Getenv("STATUSFLARE_ACCOUNT_ID"),
		keyID:     os.Getenv("STATUSFLARE_KEY_ID"),
		token:     os.Getenv("STATUSFLARE_TOKEN"),
		http:      &http.Client{},
	}

	return client, nil
}

// create new client for given account with keyID and token.
// The keyID and token are available in section
// Settings->API & Custom Workers.
//
// The account ID identify the whole account. Account might
// have multiple key IDs with tokens.
func NewClient(accountID string, keyID string, token string) *Client {
	return &Client{
		url:       "https://api.statusflare.com/",
		accountID: accountID,
		keyID:     keyID,
		token:     token,
		http:      &http.Client{},
	}
}

// all API calls are performed via this function
func (c *Client) makeAPICall(method string, endpoint string, body []byte) (*http.Response, error) {
	var err error
	var reader io.Reader

	if body != nil {
		reader = bytes.NewReader(body)
	}

	url := c.url + endpoint
	req, _ := http.NewRequest(method, url, reader)
	req.Header = map[string][]string{
		"X-Statusflare-Token":        {c.token},
		"X-Statusflare-Token-Key-Id": {c.keyID},
	}

	resp, err := c.http.Do(req)
	return resp, err
}

// helper function read the body into string
// it's usefull when you're debugging
func readBodyAsStr(resp *http.Response) string {
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(data)
}

// helps with unmarshalling response body JSON
func unmarshallResp(resp *http.Response, v interface{}) error {
	var err error

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)
	return err
}
