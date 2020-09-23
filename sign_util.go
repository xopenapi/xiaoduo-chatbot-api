package xiaoduo

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Sign(nick, question, userId, apiSecret, salt string) string {
	s := fmt.Sprintf("nick=%s&question=%s&user_id=%s%s%s", nick, question, userId, apiSecret, salt)

	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
