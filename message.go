package go_kaihei

var (
	messageUrl               = BaseApiUrl + "/message"
	listMessageUrl           = messageUrl + "/list"
	updateMessageUrl         = messageUrl + "/update"
	deleteMessageUrl         = messageUrl + "/delete"
	listReactionMessageUrl   = messageUrl + "/reaction-list"
	addReactionMessageUrl    = messageUrl + "/add-reaction"
	deleteReactionMessageUrl = messageUrl + "/delete-reaction"
)
