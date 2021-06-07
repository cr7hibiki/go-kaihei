package test

import (
	bot "github.com/drscorpio/go-kaihei"
	"os"
	"testing"
)

func TestClient_GetMe(t *testing.T) {
	c := bot.NewClient(bot.TokenTypeBot, os.Getenv("TEST_TOKEN"))
	me, err := c.GetMe()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(me.UserName)
}
