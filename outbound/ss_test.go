package outbound

import (
	"testing"
)

func TestSS(t *testing.T) {
	// rawUri := "ss://aes-256-gcm:bad5fba5-a7bc-4709-882b-e15edad16cef@ah-cmi-1000m.ikun666.club:18878#🇨🇳_CN_中国-\u003e🇸🇬_SG_新加坡"
	// rawUri := "ss://aes-128-gcm:g12sQi#ss#\u00261@183.232.170.32:20013?plugin=v2ray-plugin\u0026mode=websocket\u0026mux=undefined#🇨🇳_CN_中国-\u003e🇯🇵_JP_日本"
	rawUri := "ss://chacha20-ietf-poly1305:t0srmdxrm3xyjnvqz9ewlxb2myq7rjuv@4e168c3.h4.gladns.com:2377/?plugin=obfs-local\u0026obfs=tls\u0026obfs-host=(TG@WangCai_1)a83679f:53325#8DKJ|@Zyw_Channel"
	sso := &SShadowSocksOut{}
	sso.Parse(rawUri)
	o := sso.GetOutboundStr()
	t.Log(o)
}
