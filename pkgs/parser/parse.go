package parser

import (
	"net/url"
	"strings"

	"github.com/niubir/sing-box-parser/utils/crypt"
	"github.com/niubir/sing-box-parser/utils/logger"
)

const (
	SchemeSS        string = "ss://"
	SchemeSSR       string = "ssr://"
	SchemeTrojan    string = "trojan://"
	SchemeVless     string = "vless://"
	SchemeVmess     string = "vmess://"
	SchemeWireguard string = "wireguard://"
)

func GetVpnScheme(rawUri string) string {
	sep := "://"
	if !strings.Contains(rawUri, sep) {
		return ""
	}
	sList := strings.Split(rawUri, sep)
	return sList[0] + sep
}

func HandleQuery(rawUri string) (result string) {
	result = rawUri
	if !strings.Contains(rawUri, "?") {
		return
	}
	sList := strings.Split(rawUri, "?")
	query := sList[1]
	if strings.Contains(query, ";") && !strings.Contains(query, "&") {
		result = sList[0] + "?" + strings.ReplaceAll(sList[1], ";", "&")
	}
	return
}

func ParseRawUri(rawUri string) (result string) {
	if strings.HasPrefix(rawUri, SchemeVmess) {
		if r := crypt.DecodeBase64(strings.Split(rawUri, "://")[1]); r != "" {
			result = SchemeVmess + r
		}
		return
	}

	if strings.Contains(rawUri, "\u0026") {
		rawUri = strings.ReplaceAll(rawUri, "\u0026", "&")
	}
	rawUri, _ = url.QueryUnescape(rawUri)
	r, err := url.Parse(rawUri)
	result = rawUri
	if err != nil {
		logger.Error(err)
		return
	}

	host := r.Host
	uname := r.User.Username()
	passw, hasPassword := r.User.Password()

	if !strings.Contains(rawUri, "@") {
		if hostDecrypted := crypt.DecodeBase64(host); hostDecrypted != "" {
			result = strings.ReplaceAll(rawUri, host, hostDecrypted)
		}
	} else if uname != "" && !hasPassword && !strings.Contains(uname, "-") {
		if unameDecrypted := crypt.DecodeBase64(uname); unameDecrypted != "" {
			result = strings.ReplaceAll(rawUri, uname, unameDecrypted)
		}
	} else {
		if passwDecrypted := crypt.DecodeBase64(passw); passwDecrypted != "" {
			result = strings.ReplaceAll(rawUri, passw, passwDecrypted)
		}
	}

	if strings.Contains(result, "%") {
		result, _ = url.QueryUnescape(result)
	}
	result = HandleQuery(result)
	return
}
