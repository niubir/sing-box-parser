package outbound

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/niubir/sing-box-parser/pkgs/parser"
)

/*
http://sing-box.sagernet.org/zh/configuration/outbound/trojan/
{
  "type": "trojan",
  "tag": "trojan-out",

  "server": "127.0.0.1", # 必填
  "server_port": 1080, # 必填
  "password": "8JCsPssfgS8tiRwiMlhARg==", # 必填
  "network": "tcp",
  "tls": {},
  "multiplex": {},
  "transport": {},

  ... // 拨号字段
}

*/

var SingTrojan string = `{
	"type": "trojan",
	"tag": "trojan-out",
	"server": "127.0.0.1",
	"server_port": 1080,
	"password": "8JCsPssfgS8tiRwiMlhARg=="
}`

type STrojanOut struct {
	RawUri   string
	Parser   *parser.ParserTrojan
	outbound string
}

func (that *STrojanOut) Parse(rawUri string) {
	that.Parser = &parser.ParserTrojan{}
	that.Parser.Parse(rawUri)
}

func (that *STrojanOut) Addr() string {
	if that.Parser == nil {
		return ""
	}
	return that.Parser.GetAddr()
}

func (that *STrojanOut) Port() int {
	if that.Parser == nil {
		return 0
	}
	return that.Parser.GetPort()
}

func (that *STrojanOut) Scheme() string {
	return parser.SchemeTrojan
}

func (that *STrojanOut) GetRawUri() string {
	return that.RawUri
}

func (that *STrojanOut) getSettings() string {
	if that.Parser.Address == "" || that.Parser.Port == 0 {
		return "{}"
	}
	j := gjson.New(SingTrojan)
	j.Set("type", "trojan")
	j.Set("server", that.Parser.Address)
	j.Set("server_port", that.Parser.Port)
	j.Set("password", that.Parser.Password)
	j.Set("tag", that.Parser.Tag)
	return j.MustToJsonString()
}

func (that *STrojanOut) GetOutboundStr() string {
	if that.outbound == "" {
		settings := that.getSettings()
		if settings == "{}" {
			return ""
		}
		cnf := gjson.New(settings)
		// trojan doesn't need to prepare stream settings
		// cnf = PrepareStreamStr(cnf, that.Parser.StreamField)
		that.outbound = cnf.MustToJsonString()
	}
	return that.outbound
}
