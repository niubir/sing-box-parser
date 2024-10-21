package outbound

import (
	"testing"
)

func TestTrojan(t *testing.T) {
	// rawUri := "trojan://2a898bd8-c0d1-4f7d-a88e-831d5682a9b9@hk02.isddns.tk:65527?allowInsecure=0\u0026peer=hk02.isddns.tk\u0026sni=hk02.isddns.tk#RPD|www.zyw.asia ZYW免费节点"
	// rawUri := "trojan://da88864b-6aa5-4d18-8e36-ac809a24c571@uk.stablize.top:443?allowInsecure=1#8DKJ|@Zyw_Channel"
	rawUri := "trojan://4d706727-996f-4427-930d-60f3bd414cf9@cnamemk.ciscocdn1.live:443?type=ws\u0026sni=c2mk.ciscocdn1.live\u0026allowInsecure=1\u0026path=/rDCYQta83d0oPABKBhcIX#🇺🇸_US_美国-\u003e🇵🇱_PL_波兰"
	to := &STrojanOut{}
	to.Parse(rawUri)
	o := to.GetOutboundStr()
	t.Log(o)
}
