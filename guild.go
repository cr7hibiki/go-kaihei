package go_kaihei

import (
	"encoding/json"
)

var (
	guildUrl           = BaseApiUrl + "/guild"
	muteGuildUrl       = BaseApiUrl + "/guild-mute"
	listGuildUrl       = guildUrl + "/list"
	viewGuildUrl       = guildUrl + "/view"
	userListGuildUrl   = guildUrl + "/user-list"
	nicknameGuildUrl   = guildUrl + "/nickname"
	leaveGuildUrl      = guildUrl + "/leave"
	kickoutGuildUrl    = guildUrl + "/kickout"
	muteListGuildUrl   = muteGuildUrl + "/list"
	muteCreateGuildUrl = muteGuildUrl + "/create"
	muteDeleteGuildUrl = muteGuildUrl + "/delete"
)

type Guild struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Topic            string    `json:"topic"`
	MasterID         string    `json:"master_id"`
	Icon             string    `json:"icon"`
	NotifyType       int       `json:"notify_type"`
	Region           string    `json:"region"`
	EnableOpen       bool      `json:"enable_open"`
	OpenId           string    `json:"open_id"`
	DefaultChannelID string    `json:"default_channel_id"`
	WelcomeChannelID string    `json:"welcome_channel_id"`
	Roles            []Role    `json:"roles"`
	Channels         []Channel `json:"channels"`
}

//ListGuilds
//List all guilds which are bot joined return an array of Guild,
//And these Guilds don't include Guild.Roles and Guild.Channels properties
func (c *Client) ListGuilds() ([]Guild, error) {
	res, err := c.get(listGuildUrl)
	if err != nil {
		return nil, err
	}
	var tempData struct {
		*Status
		Data struct {
			Items []Guild  `json:"items"`
			Meta  struct { // TODO: If necessary,here needs to return Meta and sort
				Page      int `json:"page"`
				PageTotal int `json:"page_total"`
				PageSize  int `json:"page_size"`
				Total     int `json:"total"`
			} `json:"meta"`
			Sort struct {
				ID int `json:"id"`
			}
		} `json:"data"`
	}
	err = json.Unmarshal(res.Body(), &tempData)
	if err != nil {
		return nil, err
	}
	if checkResponse(tempData.Status) != nil {
		return nil, checkResponse(tempData.Status)
	}
	return tempData.Data.Items, nil
}

//ViewGuild
//Get Details of Guild Specified by guild_id
func (c *Client) ViewGuild(guildID string) (*Guild, error) {
	res, err := c.getWithParam(viewGuildUrl, "guild_id", guildID)
	if err != nil {
		return nil, err
	}
	var tempData struct {
		*Status
		Data struct {
			*Guild
			Emojis []struct { // TODO: If necessary,here needs to return Emojis and UserConfig
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"emojis"`
			UserConfig struct {
				NotifyType  interface{} `json:"notify_type"`
				NickName    string      `json:"nickname"`
				RoleIDs     []int       `json:"role_ids"`
				ChatSetting string      `json:"chat_setting"`
			} `json:"user_config"`
		} `json:"data"`
	}
	err = json.Unmarshal(res.Body(), &tempData)
	if err != nil {
		return nil, err
	}
	return tempData.Data.Guild, nil
}

//ListGuildUsers
//returns user count , online user count , offline user count ,
//and an array of User these are matched with queryParams
// queryParams:
// | ParamName       |   Type  | IsNecessary | explain
// | guild_id        | string  |      T      |  the id of specified guild
// | channel_id      | string  |      F      |  the id of specified channel
// | search 		 | string  |      F      |  query key word , will search in uer name and nickname
// | role_id		 |  int    |      F      |  query users by role
// | mobile_verified |  int    |      F      |	should be "0" or "1" ,"0" is unverified ,"1" is verified
// | active_time	 |  int    |      F      |	sort by active time "0" is order "1" is reverse order
// | joined_at       |  int    |      F      |  sort by join time "0" is order "1" is reverse order
// | page			 |  int    |      F      |  target page number
// | page_size		 |  int    |      F      |  page size
func (c Client) ListGuildUsers(queryParams map[string]string) (users []User, userCount int, userOnline int, userOffline int, err error) {
	res, err := c.getWithParams(userListGuildUrl, queryParams)
	if err != nil {
		return nil, 0, 0, 0, err
	}
	var tempData struct {
		*Status
		Data struct {
			Items        []User `json:"items"`
			UserCount    int    `json:"user_count"`
			OnlineCount  int    `json:"online_count"`
			OfflineCount int    `json:"offline_count"`
		} `json:"data"`
	}
	err = json.Unmarshal(res.Body(), &tempData)
	if err != nil {
		return nil, 0, 0, 0, err
	}
	if checkResponse(tempData.Status) != nil {
		return nil, 0, 0, 0, checkResponse(tempData.Status)
	}
	users = tempData.Data.Items
	userCount = tempData.Data.UserCount
	userOnline = tempData.Data.OnlineCount
	userOffline = tempData.Data.OfflineCount
	return
}

//ChangeNickName
//
func (c *Client) ChangeNickName() {

}
