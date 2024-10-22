package parser

import "errors"

type StreamField struct {
	Network          string
	StreamSecurity   string
	Path             string
	Host             string
	UserAgent        string
	TCPHeaderType    string
	GRPCServiceName  string
	GRPCMultiMode    string
	ServerName       string
	TLSALPN          string
	TLSAllowInsecure string
	Fingerprint      string
	RealityShortId   string
	RealitySpiderX   string
	RealityPublicKey string
	PacketEncoding   string
}

type StreamFieldSetor = func(s *StreamField) error

func NewStreamFieldNetworkSetor(v interface{}) StreamFieldSetor {
	return func(s *StreamField) error {
		tmp, ok := v.(string)
		if !ok {
			return errors.New("value not invalid")
		}
		s.Network = tmp
		return nil
	}
}
func NewStreamFieldStreamSecuritySetor(v interface{}) StreamFieldSetor {
	return func(s *StreamField) error {
		tmp, ok := v.(string)
		if !ok {
			return errors.New("value not invalid")
		}
		s.StreamSecurity = tmp
		return nil
	}
}
func NewStreamFieldPathSetor(v interface{}) StreamFieldSetor {
	return func(s *StreamField) error {
		tmp, ok := v.(string)
		if !ok {
			return errors.New("value not invalid")
		}
		s.Path = tmp
		return nil
	}
}
func NewStreamFieldHostSetor(v interface{}) StreamFieldSetor {
	return func(s *StreamField) error {
		tmp, ok := v.(string)
		if !ok {
			return errors.New("value not invalid")
		}
		s.Host = tmp
		return nil
	}
}
func NewStreamFieldUserAgentSetor(v interface{}) StreamFieldSetor {
	return func(s *StreamField) error {
		tmp, ok := v.(string)
		if !ok {
			return errors.New("value not invalid")
		}
		s.UserAgent = tmp
		return nil
	}
}
func NewStreamFieldTCPHeaderTypeSetor(v interface{}) StreamFieldSetor {
	return func(s *StreamField) error {
		tmp, ok := v.(string)
		if !ok {
			return errors.New("value not invalid")
		}
		s.TCPHeaderType = tmp
		return nil
	}
}
func NewStreamFieldGRPCServiceNameSetor(v interface{}) StreamFieldSetor {
	return func(s *StreamField) error {
		tmp, ok := v.(string)
		if !ok {
			return errors.New("value not invalid")
		}
		s.GRPCServiceName = tmp
		return nil
	}
}
func NewStreamFieldGRPCMultiModeSetor(v interface{}) StreamFieldSetor {
	return func(s *StreamField) error {
		tmp, ok := v.(string)
		if !ok {
			return errors.New("value not invalid")
		}
		s.GRPCMultiMode = tmp
		return nil
	}
}
func NewStreamFieldServerNameSetor(v interface{}) StreamFieldSetor {
	return func(s *StreamField) error {
		tmp, ok := v.(string)
		if !ok {
			return errors.New("value not invalid")
		}
		s.ServerName = tmp
		return nil
	}
}
func NewStreamFieldTLSALPNSetor(v interface{}) StreamFieldSetor {
	return func(s *StreamField) error {
		tmp, ok := v.(string)
		if !ok {
			return errors.New("value not invalid")
		}
		s.TLSALPN = tmp
		return nil
	}
}
func NewStreamFieldTLSAllowInsecureSetor(v interface{}) StreamFieldSetor {
	return func(s *StreamField) error {
		tmp, ok := v.(string)
		if !ok {
			return errors.New("value not invalid")
		}
		s.TLSAllowInsecure = tmp
		return nil
	}
}
func NewStreamFieldFingerprintSetor(v interface{}) StreamFieldSetor {
	return func(s *StreamField) error {
		tmp, ok := v.(string)
		if !ok {
			return errors.New("value not invalid")
		}
		s.Fingerprint = tmp
		return nil
	}
}
func NewStreamFieldRealityShortIdSetor(v interface{}) StreamFieldSetor {
	return func(s *StreamField) error {
		tmp, ok := v.(string)
		if !ok {
			return errors.New("value not invalid")
		}
		s.RealityShortId = tmp
		return nil
	}
}
func NewStreamFieldRealitySpiderXSetor(v interface{}) StreamFieldSetor {
	return func(s *StreamField) error {
		tmp, ok := v.(string)
		if !ok {
			return errors.New("value not invalid")
		}
		s.RealitySpiderX = tmp
		return nil
	}
}
func NewStreamFieldRealityPublicKeySetor(v interface{}) StreamFieldSetor {
	return func(s *StreamField) error {
		tmp, ok := v.(string)
		if !ok {
			return errors.New("value not invalid")
		}
		s.RealityPublicKey = tmp
		return nil
	}
}
func NewStreamFieldPacketEncodingSetor(v interface{}) StreamFieldSetor {
	return func(s *StreamField) error {
		tmp, ok := v.(string)
		if !ok {
			return errors.New("value not invalid")
		}
		s.PacketEncoding = tmp
		return nil
	}
}
