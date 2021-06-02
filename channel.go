package go_kaihei

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
	ID                   string        `json:"id"`
	Name                 string        `json:"name"`
	UserID               string        `json:"user_id"`
	GuildID              string        `json:"guild_id"`
	Topic                string        `json:"topic"`
	IsCategory           bool          `json:"is_category"`
	ParentID             string        `json:"parent_id"`
	Level                int           `json:"level"`
	SlowMode             int           `json:"slow_mode"`
	Type                 int           `json:"type"`
	PermissionOverwrites []interface{} `json:"permission_overwrites"` //TODO: it should be an some kind of array
	PermissionUsers      []interface{} `json:"permission_users"`      //TODO: it should be an some kind of array
	PermissionSync       int           `json:"permission_sync"`
}
