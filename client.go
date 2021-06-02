package go_kaihei

import (
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
