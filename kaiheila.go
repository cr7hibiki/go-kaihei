package go_kaihei

//TokenType used to set the request header
type TokenType = string

const (
	//BaseApiUrl API base URL
	BaseApiUrl = "https://www.kaiheila.cn/api/v3"

	TokenTypeBot    = TokenType("Bot")
	TokenTypeOauth2 = TokenType("Bearer")
)
