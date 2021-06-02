package go_kaihei

//TokenType used to set the request header
type TokenType = string

const (
	//BASE_API_URL API base URL
	BASE_API_URL = "https://www.kaiheila.cn/api"

	TOKEN_TYPE_BOT    = TokenType("Bot")
	TOKEN_TYPE_OAUTH2 = TokenType("Bearer")
)
