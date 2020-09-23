package xiaoduo

import (
	"context"
	"github.com/elgs/gostrgen"
	"testing"
)

func TestApi(t *testing.T) {
	cfg := NewConfiguration()
	client := NewAPIClient(cfg)

	unitId := int32(7540)
	channelId := int32(13828)
	apiSecret := "34bd1afbcb4e9a9f514fb27253e20f95"
	userId := "122222"
	nick := "小多多"
	question := "你好啊"
	salt, _ := gostrgen.RandGen(8, gostrgen.Lower | gostrgen.Upper | gostrgen.Digit, "", "")

	sign := Sign(nick, question, userId, apiSecret, salt)

	r, _, err := client.DefaultApi.MatchQuestion(context.Background(), unitId, channelId, salt, sign, userId, nick, question)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)

	question = "你是谁"
	salt, _ = gostrgen.RandGen(8, gostrgen.Lower | gostrgen.Upper | gostrgen.Digit, "", "")

	sign = Sign(nick, question, userId, apiSecret, salt)
	r, _, err = client.DefaultApi.MatchQuestion(context.Background(), unitId, channelId, salt, sign, userId, nick, question)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(r)
}