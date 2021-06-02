package go_kaihei

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
