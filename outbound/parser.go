package outbound

import (
	"errors"

	"github.com/niubir/sing-box-parser/pkgs/parser"
	"github.com/niubir/sing-box-parser/utils"
)

func NodeToOutboundStr(url string) (string, error) {
	var out string
	switch utils.ParseScheme(url) {
	case parser.SchemeSS:
		o := &SShadowSocksOut{}
		o.Parse(url)
		out = o.GetOutboundStr()
	case parser.SchemeSSR:
		o := &SShadowSocksROut{}
		o.Parse(url)
		out = o.GetOutboundStr()
	case parser.SchemeTrojan:
		o := &STrojanOut{}
		o.Parse(url)
		out = o.GetOutboundStr()
	case parser.SchemeVless:
		o := &SVlessOut{}
		o.Parse(url)
		out = o.GetOutboundStr()
	case parser.SchemeVmess:
		o := &SVmessOut{}
		o.Parse(url)
		out = o.GetOutboundStr()
	case parser.SchemeWireguard:
		o := &SWireguardOut{}
		o.Parse(url)
		out = o.GetOutboundStr()
	default:
		return "", errors.New("not support")
	}
	if out == "" {
		return "", errors.New("parse error")
	}
	return out, nil
}
