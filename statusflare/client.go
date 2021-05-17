package statusflare

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// represent the session to statusflare
type Client struct {
	apiUrl    string
	accountId string
	keyId     string
	token     string
	http      *http.Client
}

type ErrorResponse struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

// create default client where all needed data are
// read from env. variables. The required env. variables
// are:
//   - SF_ACCOUNT_ID
//   - SF_KEY_ID
//   - SF_TOKEN
//
// If none of the env. variables are set, the function returns
// you nil client and error.
func DefaultClient() (*Client, error) {
	if os.Getenv("SF_API_URL") == "" {
		return nil, fmt.Errorf("missing SF_API_URL")
	}

	if os.Getenv("SF_ACCOUNT_ID") == "" {
		return nil, fmt.Errorf("missing SF_ACCOUNT_ID")
	}

	if os.Getenv("SF_KEY_ID") == "" {
		return nil, fmt.Errorf("missing SF_KEY_ID")
	}

	if os.Getenv("SF_TOKEN") == "" {
		return nil, fmt.Errorf("missing SF_TOKEN")
	}

	client := &Client{
		apiUrl:    os.Getenv("SF_API_URL"),
		accountId: os.Getenv("SF_ACCOUNT_ID"),
		keyId:     os.Getenv("SF_KEY_ID"),
		token:     os.Getenv("SF_TOKEN"),
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
func NewClient(apiUrl string, accountId string, keyId string, token string) *Client {
	return &Client{
		apiUrl:    apiUrl,
		accountId: accountId,
		keyId:     keyId,
		token:     token,
		http:      &http.Client{},
	}
}

// all API calls are performed via this function
func (c *Client) makeAPICall(method string, endpoint string, body []byte) (*http.Response, error) {
	var err error
	var errorResponse ErrorResponse
	var reader io.Reader

	if body != nil {
		reader = bytes.NewReader(body)
	}

	url := c.apiUrl + endpoint
	req, _ := http.NewRequest(method, url, reader)
	req.Header = map[string][]string{
		"X-Statusflare-Token":        {c.token},
		"X-Statusflare-Token-Key-Id": {c.keyId},
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		err = unmarshallResp(resp, &errorResponse)
		if err != nil {
			return nil, fmt.Errorf("reading api error response failed, status code %d", resp.StatusCode)
		}
		return nil, errors.New(errorResponse.Error.Message)
	}

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
