package outbound

import (
	"testing"

	"github.com/niubir/sing-box-parser/pkgs/parser"
)

func TestVless(t *testing.T) {
	// rawUri := "vless://f0f4eabc-2747-4656-99b5-81ab6d32a8ab@172.67.33.254:443?alpn=http/1.1\u0026headerType=ws\u0026host=hct2.jensenk.cf\u0026path=/f0f4eabc-2747-4656-99b5-81ab6d32a8ab-vless\u0026security=tls\u0026sni=hct2.jensenk.cf\u0026type=ws#美国_08281722"
	// rawUri := "vless://882b8757-9244-404b-fee6-9ec7c3d8fd82@b2.v2parsin.site:17407?encryption=none\u0026security=none\u0026type=tcp\u0026headerType=http\u0026host=telewebion.com#德国_0828093"
	// rawUri := "vless://9890111b-a139-4a87-89d5-b9b18dd05e46@mci-shhproxy.ddns.net:8443?encryption=none\u0026security=tls\u0026sni=dl.SpV2ray.cfd\u0026type=grpc\u0026serviceName=@Shh_Proxy\u0026mode=gun#中国_0828245"
	rawUri := "vless://d572752d-b079-4169-a1a1-3da5721a8ab4@m2rel.siasepid.sbs:80?encryption=none\u0026allowInsecure=1\u0026security=reality\u0026sni=tgju.org\u0026fp=firefox\u0026pbk=HgrpXJzQo2liQMY9YAPq1_PuiDXNNBLx8hRyVVfUZko\u0026sid=af41f983\u0026spx=/\u0026type=grpc\u0026serviceName=@V2rayNGmat\u0026mode=multi#德国_0828096"
	// rawUri := "vless://57c3f752-3510-4908-9a84-a3490c08c022@103.140.136.37:443?encryption=none&security=tls&sni=jp01.guaziapp.me&fp=chrome&type=ws&host=jp01.guaziapp.me&path=%2F4069a771-40e3-4875-a715-853ea421b010#%F0%9F%87%AF%F0%9F%87%B5+%E6%97%A5%E6%9C%AC-01"
	vo := &SVlessOut{}
	vo.Parse(rawUri)
	setor1 := parser.NewStreamFieldUserAgentSetor("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36")
	setor2 := parser.NewStreamFieldTLSAllowInsecureSetor("1")
	setor1(vo.Parser.StreamField)
	setor2(vo.Parser.StreamField)
	o := vo.GetOutboundStr()
	t.Log("***", o)
}
