package parser

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/niubir/sing-box-parser/utils"
)

/*
vless: ['security', 'type', 'sni', 'path', 'encryption', 'headerType', 'packetEncoding', 'serviceName', 'mode', 'flow', 'alpn', 'host', 'fp', 'pbk', 'sid', 'spx']
*/

type ParserVless struct {
	Address    string
	Port       int
	UUID       string
	Encryption string
	Flow       string
	Tag        string
	*StreamField
}

func (that *ParserVless) Parse(rawUri string) {
	if u, err := url.Parse(rawUri); err == nil {
		that.Address = u.Hostname()
		that.Port, _ = strconv.Atoi(u.Port())
		that.UUID = u.User.Username()
		query := u.Query()
		that.Encryption = query.Get("encryption")
		if that.Encryption == "" {
			that.Encryption = "none"
		}
		that.Flow = query.Get("flow")
		if that.Flow == "xtls-rprx-direct-udp443" {
			that.Flow = "xtls-rprx-vision-udp443"
		}
		that.Tag = utils.GenTag(u.Fragment)

		that.StreamField = &StreamField{
			Network:          query.Get("type"),
			StreamSecurity:   query.Get("security"),
			Path:             query.Get("path"),
			Host:             query.Get("host"),
			UserAgent:        query.Get("userAgent"),
			GRPCServiceName:  query.Get("serviceName"),
			GRPCMultiMode:    query.Get("mode"),
			ServerName:       query.Get("sni"),
			TLSALPN:          query.Get("alpn"),
			Fingerprint:      query.Get("fp"),
			RealityShortId:   query.Get("sid"),
			RealitySpiderX:   query.Get("spx"),
			RealityPublicKey: query.Get("pbk"),
			PacketEncoding:   query.Get("packetEncoding"),
			TCPHeaderType:    query.Get("headerType"),
			TLSAllowInsecure: query.Get("allowInsecure"),
		}
	}
}

func (that *ParserVless) GetAddr() string {
	return that.Address
}

func (that *ParserVless) GetPort() int {
	return that.Port
}

func (that *ParserVless) Show() {
	fmt.Printf("addr: %s, port: %v, uuid: %s\n",
		that.Address,
		that.Port,
		that.UUID)
}
