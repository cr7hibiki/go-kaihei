package go_kaihei

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

//Client a basic Kaiheila API client.
//Use NewClient method to create a new client.
type Client struct {
	token     string
	tokenType TokenType
	client    *resty.Client
}

//NewClient
//Create a new Client with specified TokenType and your own Token
func NewClient(tokenType TokenType, token string) *Client {
	c := resty.New()
	c.SetHeader("Content-Type", "application/json")
	c.SetHeader("Authorization", fmt.Sprintf("%s %s", tokenType, token))
	return &Client{
		client:    c,
		token:     token,
		tokenType: tokenType,
	}
}

func (c *Client) get(url string) (*resty.Response, error) {
	return c.client.R().Get(url)
}

func (c *Client) getWithParam(url string, key string, value string) (*resty.Response, error) {
	return c.client.R().SetQueryParam(key, value).Get(url)
}

func (c *Client) getWithParams(url string, queryParams map[string]string) (*resty.Response, error) {
	return c.client.R().SetQueryParams(queryParams).Get(url)
}

func (c *Client) post(url string, body interface{}) (*resty.Response, error) {
	return c.client.R().SetBody(body).Post(url)
}

func (c *Client) PostFile(url string, param string, path string) (*resty.Response, error) {
	return c.client.R().SetFile(param, path).Post(url)
}

func checkResponse(status *Status) error {
	if status.Code != 0 {
		return errors.New(fmt.Sprintf("Error code %d : %s", status.Code, status.Msg))
	} else {
		return nil
	}
}
