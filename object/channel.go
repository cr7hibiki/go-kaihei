package object

type Channel struct {
	ID                   string      `json:"id"`
	Name                 string      `json:"name"`
	UserID               string      `json:"user_id"`
	GuildID              string      `json:"guild_id"`
	Topic                string      `json:"topic"`
	IsCategory           bool        `json:"is_category"`
	ParentID             string      `json:"parent_id"`
	Level                int         `json:"level"`
	SlowMode             int         `json:"slow_mode"`
	Type                 int         `json:"type"`
	PermissionOverwrites string      `json:"permission_overwrites"`
	PermissionUsers      interface{} `json:"permission_users"` //TODO: it should be an some kind of array
	PermissionSync       int         `json:"permission_sync"`
}
