package go_kaihei

type Role struct {
	RoleID      int    `json:"role_id"`
	Name        string `json:"name"`
	Color       int    `json:"color"`
	Position    int    `json:"position"`
	Hoist       int    `json:"hoist"`
	Mentionable int    `json:"mentionable"`
	Permissions int    `json:"permissions"`
}
