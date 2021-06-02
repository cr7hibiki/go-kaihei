package go_kaihei

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
