package parser

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/niubir/sing-box-parser/utils"
	"github.com/niubir/sing-box-parser/utils/logger"
)

/*
trojan: ['allowInsecure', 'peer', 'sni', 'type', 'path', 'security', 'headerType']
*/

type ParserTrojan struct {
	Address  string
	Port     int
	Password string
	Tag      string
	*StreamField
}

func (that *ParserTrojan) Parse(rawUri string) {
	if u, err := url.Parse(rawUri); err == nil {
		that.Address = u.Hostname()
		that.Port, _ = strconv.Atoi(u.Port())
		that.Password = u.User.Username()
		that.Tag = utils.GenTag(u.Fragment)
		query := u.Query()
		that.StreamField = &StreamField{
			Network:          query.Get("type"),
			Host:             query.Get("peer"),
			UserAgent:        query.Get("userAgent"),
			Path:             query.Get("path"),
			StreamSecurity:   query.Get("security"),
			ServerName:       query.Get("sni"),
			TCPHeaderType:    query.Get("headerType"),
			TLSAllowInsecure: query.Get("allowInsecure"),
		}
	} else {
		logger.Error(err)
		fmt.Println(rawUri)
		return
	}

	if that.StreamField.TLSAllowInsecure != "" && that.StreamField.ServerName != "" {
		if that.Network == "" {
			that.Network = "tcp"
		}
		if that.StreamSecurity == "" {
			that.StreamSecurity = "tls"
		}
		if that.Host == "" {
			that.Host = that.ServerName
		}
		if that.ServerName == "" {
			that.ServerName = that.Address
		}
	}
}

func (that *ParserTrojan) GetAddr() string {
	return that.Address
}

func (that *ParserTrojan) GetPort() int {
	return that.Port
}

func (that *ParserTrojan) Show() {
	fmt.Printf("addr: %s, port: %v, password: %s\n",
		that.Address,
		that.Port,
		that.Password)
}
