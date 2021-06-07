package go_kaihei

import "encoding/json"

var (
	channelUrl           = BaseApiUrl + "/channel"
	channelRoleUrl       = BaseApiUrl + "/channel-role"
	listChannelUrl       = channelUrl + "/list"
	viewChannelUrl       = channelUrl + "/view"
	createChannelUrl     = channelUrl + "/create"
	deleteChannelUrl     = channelUrl + "/delete"
	moveUserChannelUrl   = channelUrl + "/move-user"
	indexChannelRoleUrl  = channelRoleUrl + "/index"
	createChannelRoleUrl = channelRoleUrl + "/create"
	updateChannelRoleUrl = channelRoleUrl + "/update"
	deleteChannelRoleUrl = channelRoleUrl + "/delete"
)

type Channel struct {
	ID                   string                `json:"id"`
	Name                 string                `json:"name"`
	MasterID             string                `json:"master_id"`
	GuildID              string                `json:"guild_id"`
	Topic                string                `json:"topic"`
	LimitAmount          int                   `json:"limit_amount"`
	IsCategory           bool                  `json:"is_category"`
	ParentID             string                `json:"parent_id"`
	Level                int                   `json:"level"`
	SlowMode             int                   `json:"slow_mode"`
	Type                 int                   `json:"type"`
	ServerURL            string                `json:"server_url"`
	PermissionOverwrites []PermissionOverwrite `json:"permission_overwrites"`
	PermissionUsers      []PermissionUser      `json:"permission_users"`
	PermissionSync       int                   `json:"permission_sync"`
}

type PermissionOverwrite struct {
	RoleID int `json:"role_id"`
	*PermissionSet
}

type PermissionUser struct {
	User User `json:"user"`
	*PermissionSet
}

type PermissionSet struct {
	Allow int `json:"allow"`
	Deny  int `json:"deny"`
}

//ListChannel
//return an array of Channel these are belong to Guild
//these Channels don't have properties of
func (c *Client) ListChannel(guildID string) ([]Channel, error) {
	res, err := c.getWithParam(listChannelUrl, "target_id", guildID)
	if err != nil {
		return nil, err
	}
	var tempData struct {
		*Status
		Data struct {
			Items []Channel `json:"items"`
			Meta  struct {  // TODO: If necessary,here needs to return Meta and sort
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

//ViewChannel
//
func (c *Client) ViewChannel(channelID string) (*Channel, error) {
	res, err := c.getWithParam(viewChannelUrl, "target_id", channelID)
	if err != nil {
		return nil, err
	}
	var tempData struct {
		*Status
		Data *Channel `json:"data"`
	}
	err = json.Unmarshal(res.Body(), &tempData)
	if err != nil {
		return nil, err
	}
	if checkResponse(tempData.Status) != nil {
		return nil, checkResponse(tempData.Status)
	}
	return tempData.Data, nil
}

//CreateChannel
//TODO: finish params comments.
func (c *Client) CreateChannel(params map[string]string) (*Channel, error) {
	res, err := c.post(createChannelUrl, params)
	if err != nil {
		return nil, err
	}
	var tempData struct {
		*Status
		Data *Channel `json:"data"`
	}
	err = json.Unmarshal(res.Body(), &tempData)
	if err != nil {
		return nil, err
	}
	if checkResponse(tempData.Status) != nil {
		return nil, checkResponse(tempData.Status)
	}
	return tempData.Data, nil
}

func (c *Client) DeleteChannel(channelID string) error {
	res, err := c.post(deleteChannelUrl, map[string]string{"channel_id": channelID})
	if err != nil {
		return err
	}
	tempData := &Status{}
	err = json.Unmarshal(res.Body(), tempData)
	if err != nil {
		return err
	}
	if checkResponse(tempData) != nil {
		return checkResponse(tempData)
	}
	return nil
}

//MoveUserChannel
//
func (c *Client) MoveUserChannel(channelID string, userIDs []string) error {
	data := map[string]interface{}{
		"target_id": channelID,
		"user_ids":  userIDs,
	}
	res, err := c.post(moveUserChannelUrl, data)
	if err != nil {
		return nil
	}
	var tempData *Status
	err = json.Unmarshal(res.Body(), tempData)
	if err != nil {
		return err
	}
	if checkResponse(tempData) != nil {
		return checkResponse(tempData)
	}
	return nil
}

//IndexChannelRole
//return array of PermissionOverwrite and PermissionUser
func (c *Client) IndexChannelRole(channelID string) (PermissionOverwrites []PermissionOverwrite, PermissionUsers []PermissionUser, PermissionSync int, err error) {
	res, err := c.getWithParam(indexChannelRoleUrl, "channel_id", channelID)
	if err != nil {
		return
	}
	var tempData struct {
		*Status
		Data struct {
			PermissionOverwrites []PermissionOverwrite `json:"permission_overwrites"`
			PermissionUsers      []PermissionUser      `json:"permission_users"`
			PermissionSync       int                   `json:"permission_sync"`
		} `json:"data"`
	}
	err = json.Unmarshal(res.Body(), &tempData)
	if err != nil {
		return
	}
	if checkResponse(tempData.Status) != nil {
		err = checkResponse(tempData.Status)
		return
	}
	PermissionOverwrites = tempData.Data.PermissionOverwrites
	PermissionUsers = tempData.Data.PermissionUsers
	PermissionSync = tempData.Data.PermissionSync
	return
}

//CreateChannelRole
//TODO: complete comments
func (c *Client) CreateChannelRole(params ...string) error {
	return nil

}

func (c *Client) UpdateChannelRole(params ...string) ([]PermissionOverwrite, error) {
	return nil, nil
}

func (c *Client) DeleteChannelRole(params ...string) error {
	return nil
}
