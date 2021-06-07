package go_kaihei

import (
	"encoding/json"
	"errors"
	"strconv"
)

var (
	gatewayUrl      = BaseApiUrl + "/gateway"
	indexGatewayUrl = gatewayUrl + "/index"
)

//GetGateway
//param compress must be "1" or "0"
//if compress is empty it will be default value "1"
// "1" represents do compress when communicating with WebSocket or WebHook
// "0" represents no compress
func (c *Client) GetGateway(compress ...int) (url string, err error) {
	if len(compress) > 1 {
		return "", errors.New("quantity of params compress must be 1 or 0")
	}
	if len(compress) == 0 {
		compress = []int{1}
	}
	res, err := c.getWithParam(indexGatewayUrl, "compress", strconv.Itoa(compress[0]))
	if err != nil {
		return "", err
	}
	var tempData struct {
		*Status
		Data struct {
			Url string `json:"url"`
		} `json:"data"`
	}
	err = json.Unmarshal(res.Body(), &tempData)
	if err != nil {
		return "", err
	}
	if checkResponse(tempData.Status) != nil {
		return "", checkResponse(tempData.Status)
	}
	url = tempData.Data.Url
	return
}
