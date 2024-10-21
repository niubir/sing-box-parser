package outbound

import "testing"

func TestSSR(t *testing.T) {
	// rawUri := "ssr://94.23.116.190:443:origin:aes-256-ctr:tls1.2_ticket_authSG93ZHlCeXBhc3NlcjIwMjI=remarks=MTJ8UmxKZmMzQmxaV1J1YjJSbFh6QXdNalUlM0Q=\u0026obfsparam=VG05dVpRJTNEJTNE\u0026protoparam=VG05dVpRJTNEJTNE"
	rawUri := "ssr://94.23.116.190:443:origin:aes-256-ctr:tls1.2_ticket_auth:SG93ZHlCeXBhc3NlcjIwMjI=/?obfsparam=Tm9uJSXvv70lJe+/vVxceDFm\u0026protoparam=Tm9uJSXvv70lJe+/vVxceDFm\u0026remarks=5rOV5Zu9XzA4MjgwMDk\u0026group="
	sso := &SShadowSocksROut{}
	sso.Parse(rawUri)
	o := sso.GetOutboundStr()
	t.Log(o)
}
