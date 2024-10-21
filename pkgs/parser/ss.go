package parser

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/niubir/sing-box-parser/utils"
	"github.com/niubir/sing-box-parser/utils/crypt"
)

var SSMethod map[string]struct{} = map[string]struct{}{
	"2022-blake3-aes-128-gcm":       {},
	"2022-blake3-aes-256-gcm":       {},
	"2022-blake3-chacha20-poly1305": {},
	"none":                          {},
	"aes-128-gcm":                   {},
	"aes-192-gcm":                   {},
	"aes-256-gcm":                   {},
	"chacha20-ietf-poly1305":        {},
	"xchacha20-ietf-poly1305":       {},
	"aes-128-ctr":                   {},
	"aes-192-ctr":                   {},
	"aes-256-ctr":                   {},
	"aes-128-cfb":                   {},
	"aes-192-cfb":                   {},
	"aes-256-cfb":                   {},
	"rc4-md5":                       {},
	"chacha20-ietf":                 {},
	"xchacha20":                     {},
}

/*
shadowsocks: ['plugin', 'obfs', 'obfs-host', 'mode', 'path', 'mux', 'host']
*/

type ParserSS struct {
	Address  string
	Port     int
	Method   string
	Password string
	Host     string
	Mode     string
	Mux      string
	Path     string
	Plugin   string
	OBFS     string
	OBFSHost string
	Tag      string
	*StreamField
}

func (that *ParserSS) Parse(rawUri string) {
	rawUri = that.handleSS(rawUri)
	if u, err := url.Parse(rawUri); err == nil {
		that.StreamField = &StreamField{}
		that.Address = u.Hostname()
		that.Port, _ = strconv.Atoi(u.Port())
		that.Method = u.User.Username()
		if that.Method == "rc4" {
			that.Method = "rc4-md5"
		}
		if _, ok := SSMethod[that.Method]; !ok {
			that.Method = "none"
		}
		that.Password, _ = u.User.Password()

		query := u.Query()
		that.Host = query.Get("host")
		that.Mode = query.Get("mode")
		that.Mux = query.Get("mux")
		that.Path = query.Get("path")
		that.Plugin = query.Get("plugin")
		that.OBFS = query.Get("obfs")
		that.OBFSHost = query.Get("obfs-host")
		that.Tag = utils.GenTag(u.Fragment)
	}
}

func (that *ParserSS) handleSS(rawUri string) string {
	rawUri = strings.ReplaceAll(rawUri, "#ss#\u00261@", "@")
	if u, err := url.Parse(rawUri); err == nil {
		if crypt.IsValidBase64(u.User.Username()) {
			tmp := crypt.DecodeBase64(u.User.Username())
			rawUri = strings.ReplaceAll(rawUri, u.User.Username(), tmp)
		}
	}
	return rawUri
}

func (that *ParserSS) GetAddr() string {
	return that.Address
}

func (that *ParserSS) GetPort() int {
	return that.Port
}

func (that *ParserSS) Show() {
	fmt.Printf("addr: %s, port: %d, method: %s, password: %s\n",
		that.Address,
		that.Port,
		that.Method,
		that.Password)
}
