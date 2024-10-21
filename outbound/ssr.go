package outbound

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/niubir/sing-box-parser/pkgs/parser"
)

/*
http://sing-box.sagernet.org/zh/configuration/outbound/shadowsocksr/#_3
{
  "type": "shadowsocksr",
  "tag": "ssr-out",

  "server": "127.0.0.1", # 必填
  "server_port": 1080, # 必填
  "method": "aes-128-cfb", # 必填
  "password": "8JCsPssfgS8tiRwiMlhARg==", # 必填
  "obfs": "plain",
  "obfs_param": "",
  "protocol": "origin",
  "protocol_param": "",
  "network": "udp",

  ... // 拨号字段
}

Protocal:
origin
verify_sha1
auth_sha1_v4
auth_aes128_md5
auth_aes128_sha1
auth_chain_a
auth_chain_b

OBFS:
plain
http_simple
http_post
random_head
tls1.2_ticket_auth

*/

var SingSSR string = `{
	"type": "shadowsocksr",
	"tag": "ssr-out",
	"server": "127.0.0.1",
	"server_port": 1080,
	"method": "aes-128-cfb",
	"password": "8JCsPssfgS8tiRwiMlhARg=="
}`

type SShadowSocksROut struct {
	RawUri   string
	Parser   *parser.ParserSSR
	outbound string
}

func (that *SShadowSocksROut) Parse(rawUri string) {
	that.Parser = &parser.ParserSSR{}
	that.Parser.Parse(rawUri)
}

func (that *SShadowSocksROut) Addr() string {
	if that.Parser == nil {
		return ""
	}
	return that.Parser.GetAddr()
}

func (that *SShadowSocksROut) Port() int {
	if that.Parser == nil {
		return 0
	}
	return that.Parser.GetPort()
}

func (that *SShadowSocksROut) Scheme() string {
	return parser.SchemeSSR
}

func (that *SShadowSocksROut) GetRawUri() string {
	return that.RawUri
}

func (that *SShadowSocksROut) getSettings() string {
	if that.Parser.Address == "" || that.Parser.Port == 0 {
		return "{}"
	}
	j := gjson.New(SingSS)
	j.Set("type", "shadowsocksr")
	j.Set("server", that.Parser.Address)
	j.Set("server_port", that.Parser.Port)
	j.Set("method", that.Parser.Method)
	j.Set("password", that.Parser.Password)
	j.Set("tag", that.Parser.Tag)

	if that.Parser.Proto != "" {
		j.Set("protocol", that.Parser.Proto)
	}

	if that.Parser.ProtoParam != "" {
		j.Set("protocol_param", that.Parser.ProtoParam)
	}

	if that.Parser.OBFS != "" {
		j.Set("obfs", that.Parser.OBFS)
	}

	if that.Parser.OBFSParam != "" {
		j.Set("protocol_param", that.Parser.OBFSParam)
	}
	return j.MustToJsonString()
}

func (that *SShadowSocksROut) GetOutboundStr() string {
	if that.outbound == "" {
		settings := that.getSettings()
		if settings == "{}" {
			return ""
		}
		cnf := gjson.New(settings)
		// cnf = PrepareStreamStr(cnf, that.Parser.StreamField)
		that.outbound = cnf.MustToJsonString()
	}
	return that.outbound
}
