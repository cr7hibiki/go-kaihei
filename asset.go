package go_kaihei

import "encoding/json"

var (
	assetUrl       = BaseApiUrl + "/asset"
	createAssetUrl = assetUrl + "/create"
)

func (c *Client) CreateAsset(param string, path string) (url string, err error) {
	res, err := c.PostFile(createAssetUrl, param, path)
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
