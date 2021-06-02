package go_kaihei

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
