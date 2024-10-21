package outbound

import "testing"

func TestVmess(t *testing.T) {
	// rawUri := "vmess://{\"v\": \"2\", \"ps\": \"13|西班牙 02 | 1x ES\", \"add\": \"2d3e6s01.mcfront.xyz\", \"port\": \"31884\", \"aid\": 0, \"scy\": \"auto\", \"net\": \"tcp\", \"type\": \"none\", \"tls\": \"tls\", \"id\": \"82a934c7-d98d-4e08-b63f-827b132d42ac\", \"sni\": \"es04.lovemc.xyz\"}"
	// rawUri := "vmess://{\"add\":\"bobbykotick.rip\",\"host\":\"Kansas.bobbykotick.rip\",\"sni\":\"Kansas.bobbykotick.rip\",\"id\":\"D213ED80-199B-4A01-9D62-BBCBA9C16226\",\"net\":\"ws\",\"path\":\"\\/speedtest\",\"port\":\"443\",\"ps\":\"GetAFreeNode.com-Kansas\",\"tls\":\"tls\",\"fp\":\"android\",\"alpn\":\"h2,http\\/1.1\",\"v\":2,\"aid\":0,\"type\":\"none\"}"
	rawUri := "vmess://ew0KICAidiI6ICIyIiwNCiAgInBzIjogIvCfh6jwn4ez5Y+w5rm+QjAzfOWOn+eUn0lQLeaOqOiNkCIsDQogICJhZGQiOiAiMTgzLjI0MC4xNzkuMjA3IiwNCiAgInBvcnQiOiAiMzAxMTgiLA0KICAiaWQiOiAiNTRmYzI1MDQtZjhkMS00NmY5LTlmZjUtZjQxOTk1M2QwMmZiIiwNCiAgImFpZCI6ICIwIiwNCiAgInNjeSI6ICJhdXRvIiwNCiAgIm5ldCI6ICJ0Y3AiLA0KICAidHlwZSI6ICJub25lIiwNCiAgImhvc3QiOiAiIiwNCiAgInBhdGgiOiAiIiwNCiAgInRscyI6ICIiLA0KICAic25pIjogIiIsDQogICJhbHBuIjogIiIsDQogICJmcCI6ICIiDQp9"
	vo := &SVmessOut{}
	vo.Parse(rawUri)
	o := vo.GetOutboundStr()
	t.Log(o)
}
