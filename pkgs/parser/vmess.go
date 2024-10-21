package parser

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/niubir/sing-box-parser/utils"
	"github.com/niubir/sing-box-parser/utils/crypt"
)

/*
vmess: ['v', 'ps', 'add', 'port', 'aid', 'scy', 'net', 'type', 'tls', 'id', 'sni', 'host', 'path', 'alpn', 'security', 'skip-cert-verify', 'fp', 'test_name', 'serverPort', 'nation']
*/

type ParserVmess struct {
	Address        string
	Port           int
	UUID           string
	Security       string
	AID            string
	Nation         string
	Path           string
	PS             string
	ServerPort     string
	SkipCertVerify bool
	TestName       string
	V              string
	Tag            string
	*StreamField
}

func (that *ParserVmess) Parse(rawUri string) {
	r := strings.ReplaceAll(rawUri, SchemeVmess, "")
	if crypt.IsValidBase64(r) {
		r = crypt.DecodeBase64(r)
	}
	j := gjson.New(r)
	if j == nil {
		return
	}
	that.Address = j.Get("add").String()
	if !strings.Contains(that.Address, ".") {
		that.Address = ""
		return
	}
	that.Port = j.Get("port").Int()
	that.UUID = j.Get("id").String()
	that.AID = j.Get("aid").String()
	that.Security = j.Get("security").String()
	that.Security = j.Get("security").String()
	if that.Security == "" {
		that.Security = j.Get("scy").String()
	}

	that.Nation = j.Get("nation").String()
	that.PS = j.Get("ps").String()
	that.ServerPort = j.Get("serverPort").String()
	that.SkipCertVerify = j.Get("skip-cert-verify").Bool()
	that.TestName = j.Get("test_name").String()
	that.V = j.Get("v").String()

	that.StreamField = &StreamField{}
	that.StreamField.Network = j.Get("net").String()
	that.StreamField.StreamSecurity = j.Get("tls").String()
	that.StreamField.Path = j.Get("path").String()
	that.StreamField.Host = j.Get("host").String()
	// that.StreamField.GRPCServiceName = j.GetString("serviceName")
	// that.StreamField.GRPCMultiMode = j.GetString("mode")
	that.StreamField.ServerName = j.Get("sni").String()
	that.StreamField.TCPHeaderType = j.Get("type").String()
	that.StreamField.TLSALPN = j.Get("alpn").String()
	that.StreamField.Fingerprint = j.Get("fp").String()

	// that.StreamField.RealityShortId = j.GetString("sid")
	// that.StreamField.RealitySpiderX = j.GetString("spx")
	// that.StreamField.RealityPublicKey = j.GetString("pbk")

	that.Tag = utils.GenTag(j.Get("ps").String())
}

func (that *ParserVmess) GetAddr() string {
	return that.Address
}

func (that *ParserVmess) GetPort() int {
	return that.Port
}

func (that *ParserVmess) Show() {
	fmt.Printf("addr: %s, port: %v, uuid: %s, net: %s", that.Address, that.Port, that.UUID, that.Network)
}
