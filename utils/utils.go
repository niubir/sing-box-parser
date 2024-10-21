package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/niubir/sing-box-parser/utils/crypt"
)

const (
	tagCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func SetJsonObjectByString(key, value string, gJSON *gjson.Json) (newGJSON *gjson.Json) {
	if gJSON == nil {
		return
	}
	tempValue := "OOXXOOXX"
	gJSON.Set(key, tempValue)
	result := strings.ReplaceAll(gJSON.MustToJsonString(), fmt.Sprintf(`"%s"`, tempValue), value)
	return gjson.New(result)
}

func GenTag(tag string) string {
	tag = strings.TrimSpace(tag)
	if tag != "" {
		if tmp := crypt.DecodeBase64(tag); tmp != "" {
			tag = tmp
		}
		return tag
	}
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	b := make([]byte, 10)
	for i := range b {
		b[i] = tagCharset[r.Intn(len(tagCharset))]
	}
	return fmt.Sprintf("PROXY_%s", string(b))
}
