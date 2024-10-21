package outbound

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/niubir/sing-box-parser/pkgs/parser"
)

/*
http://sing-box.sagernet.org/zh/configuration/outbound/vless/

{
  "type": "vless",
  "tag": "vless-out",

  "server": "127.0.0.1", # 必填
  "server_port": 1080, # 必填
  "uuid": "bf000d23-0752-40b4-affe-68f7707a9661", # 必填
  "flow": "xtls-rprx-vision",
  "network": "tcp",
  "tls": {},
  "packet_encoding": "",
  "transport": {},

  ... // 拨号字段
}

*/

var SingVless string = `{
	"type": "vless",
	"tag": "vless-out",
	"server": "127.0.0.1",
	"server_port": 1080,
	"uuid": "bf000d23-0752-40b4-affe-68f7707a9661",
	"flow": ""
}`

type SVlessOut struct {
	RawUri   string
	Parser   *parser.ParserVless
	outbound string
}

func (that *SVlessOut) Parse(rawUri string) {
	that.Parser = &parser.ParserVless{}
	that.Parser.Parse(rawUri)
}

func (that *SVlessOut) Addr() string {
	if that.Parser == nil {
		return ""
	}
	return that.Parser.GetAddr()
}

func (that *SVlessOut) Port() int {
	if that.Parser == nil {
		return 0
	}
	return that.Parser.GetPort()
}

func (that *SVlessOut) Scheme() string {
	return parser.SchemeVless
}

func (that *SVlessOut) GetRawUri() string {
	return that.RawUri
}

func (that *SVlessOut) getSettings() string {
	if that.Parser.Address == "" || that.Parser.Port == 0 {
		return "{}"
	}
	j := gjson.New(SingVless)
	j.Set("type", "vless")
	j.Set("server", that.Parser.Address)
	j.Set("server_port", that.Parser.Port)
	j.Set("uuid", that.Parser.UUID)
	j.Set("flow", that.Parser.Flow)
	if that.Parser.PacketEncoding != "" {
		j.Set("packet_encoding", that.Parser.PacketEncoding)
	}
	j.Set("tag", that.Parser.Tag)
	return j.MustToJsonString()
}

func (that *SVlessOut) GetOutboundStr() string {
	if that.outbound == "" {
		settings := that.getSettings()
		if settings == "{}" {
			return ""
		}
		cnf := gjson.New(settings)
		cnf = PrepareStreamStr(cnf, that.Parser.StreamField)
		that.outbound = cnf.MustToJsonString()
	}
	return that.outbound
}
