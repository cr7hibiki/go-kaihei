package go_kaihei

import "encoding/json"

var (
	userUrl = BaseApiUrl + "/user"
	meUrl   = userUrl + "/me"
)

type User struct {
	ID             string `json:"id"`
	UserName       string `json:"username"`
	IdentifyNum    string `json:"identify_num"`
	Online         bool   `json:"online"`
	Status         int    `json:"status"`
	Avatar         string `json:"avatar"`
	Bot            bool   `json:"bot"`
	MobileVerified bool   `json:"mobile_verified"`
	System         bool   `json:"system"`
	MobilePrefix   string `json:"mobile_prefix"`
	Mobile         string `json:"mobile"`
	InviteCount    int    `json:"invite_count"`
	NickName       string `json:"nickname"`
	Roles          []Role `json:"roles"`
}

func (c *Client) GetMe() (*User, error) {
	res, err := c.get(meUrl)
	if err != nil {
		return nil, err
	}
	var tempData struct {
		*Status
		Data User
	}
	err = json.Unmarshal(res.Body(), &tempData)
	if err != nil {
		return nil, err
	}
	if checkResponse(tempData.Status) != nil {
		return nil, checkResponse(tempData.Status)
	}
	return &tempData.Data, nil
}
