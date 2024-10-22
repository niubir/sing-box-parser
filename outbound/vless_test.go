package outbound

import (
	"testing"
)

func TestVless(t *testing.T) {
	// rawUri := "vless://f0f4eabc-2747-4656-99b5-81ab6d32a8ab@172.67.33.254:443?alpn=http/1.1\u0026headerType=ws\u0026host=hct2.jensenk.cf\u0026path=/f0f4eabc-2747-4656-99b5-81ab6d32a8ab-vless\u0026security=tls\u0026sni=hct2.jensenk.cf\u0026type=ws#美国_08281722"
	// rawUri := "vless://882b8757-9244-404b-fee6-9ec7c3d8fd82@b2.v2parsin.site:17407?encryption=none\u0026security=none\u0026type=tcp\u0026headerType=http\u0026host=telewebion.com#德国_0828093"
	// rawUri := "vless://9890111b-a139-4a87-89d5-b9b18dd05e46@mci-shhproxy.ddns.net:8443?encryption=none\u0026security=tls\u0026sni=dl.SpV2ray.cfd\u0026type=grpc\u0026serviceName=@Shh_Proxy\u0026mode=gun#中国_0828245"
	rawUri := "vless://d572752d-b079-4169-a1a1-3da5721a8ab4@m2rel.siasepid.sbs:80?encryption=none\u0026allowInsecure=1\u0026security=reality\u0026sni=tgju.org\u0026fp=firefox\u0026pbk=HgrpXJzQo2liQMY9YAPq1_PuiDXNNBLx8hRyVVfUZko\u0026sid=af41f983\u0026spx=/\u0026type=grpc\u0026serviceName=@V2rayNGmat\u0026mode=multi#德国_0828096"
	vo := &SVlessOut{}
	vo.Parse(rawUri)
	o := vo.GetOutboundStr()
	t.Log("***", o)
}
