package go_kaihei

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

//Client a basic Kaiheila API client.
//Use NewClient method to create a new client.
type Client struct {
	tocken     string
	tockenType TokenType
	client     *resty.Client
}

//NewClient
func NewClient(tockenType TokenType, tocken string) *Client {
	c := resty.New()
	c.SetHeader("Content-Type", "application/json")
	c.SetAuthToken(fmt.Sprintf("%s %s", tockenType, tocken))
	return &Client{
		client:     c,
		tocken:     tocken,
		tockenType: tockenType,
	}
}

func (c *Client) get(url string) (*resty.Response, error) {
	return c.client.R().Get(url)
}

func (c *Client) getWithParam(url string, key string, value string) (*resty.Response, error) {
	return c.client.R().SetQueryParam(key, value).Get(url)
}

func (c Client) getWithParams(url string, queryParams map[string]string) (*resty.Response, error) {
	return c.client.R().SetQueryParams(queryParams).Get(url)
}

func checkResponse(status *Status) error {
	if status.Code != 0 {
		return errors.New(fmt.Sprintf("Error code %d : %s", status.Code, status.Msg))
	} else {
		return nil
	}
}
