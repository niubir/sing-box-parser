package outbound

import (
	"errors"

	"github.com/niubir/sing-box-parser/pkgs/parser"
	"github.com/niubir/sing-box-parser/utils"
)

type nodeToOutboundConfig struct {
	streamSetors []parser.StreamFieldSetor
}

type Option func(*nodeToOutboundConfig)

func WithStreamSetors(streamSetors ...parser.StreamFieldSetor) func(*nodeToOutboundConfig) {
	return func(c *nodeToOutboundConfig) {
		c.streamSetors = streamSetors
	}
}

func NodeToOutboundStr(url string, options ...Option) (string, error) {
	var c nodeToOutboundConfig
	for _, option := range options {
		option(&c)
	}

	var out string
	switch utils.ParseScheme(url) {
	case parser.SchemeSS:
		o := &SShadowSocksOut{}
		o.Parse(url)
		for _, streamSetor := range c.streamSetors {
			if err := streamSetor(o.Parser.StreamField); err != nil {
				return "", err
			}
		}
		out = o.GetOutboundStr()
	case parser.SchemeSSR:
		o := &SShadowSocksROut{}
		o.Parse(url)
		for _, streamSetor := range c.streamSetors {
			if err := streamSetor(o.Parser.StreamField); err != nil {
				return "", err
			}
		}
		out = o.GetOutboundStr()
	case parser.SchemeTrojan:
		o := &STrojanOut{}
		o.Parse(url)
		for _, streamSetor := range c.streamSetors {
			if err := streamSetor(o.Parser.StreamField); err != nil {
				return "", err
			}
		}
		out = o.GetOutboundStr()
	case parser.SchemeVless:
		o := &SVlessOut{}
		o.Parse(url)
		for _, streamSetor := range c.streamSetors {
			if err := streamSetor(o.Parser.StreamField); err != nil {
				return "", err
			}
		}
		out = o.GetOutboundStr()
	case parser.SchemeVmess:
		o := &SVmessOut{}
		o.Parse(url)
		for _, streamSetor := range c.streamSetors {
			if err := streamSetor(o.Parser.StreamField); err != nil {
				return "", err
			}
		}
		out = o.GetOutboundStr()
	case parser.SchemeWireguard:
		o := &SWireguardOut{}
		o.Parse(url)
		for _, streamSetor := range c.streamSetors {
			if err := streamSetor(o.Parser.StreamField); err != nil {
				return "", err
			}
		}
		out = o.GetOutboundStr()
	default:
		return "", errors.New("not support")
	}
	if out == "" {
		return "", errors.New("parse error")
	}
	return out, nil
}
