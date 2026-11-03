//ff:func feature=live type=helper control=iteration dimension=1
//ff:what TestIsDocPath — 문서·코덱·목록의 버전과 versions.geul은 문서 주소, <Issuer>/hashes/와 .geul 아닌 경로는 아니다
package tally

import "testing"

func TestIsDocPath(t *testing.T) {
	for _, p := range []string{
		"0x00000001/docs/0x0000000000000001/versions.geul",
		"0x00000001/docs/0x0000000000000001/0x00000001.geul",
		"0x00000000/codecs/0x00000140/0x00000001.geul",
		"0x00000001/docs/0x00000001.geul",
		"0x00000000/codecs/versions.geul",
		"issuers/0x00000001.geul",
		"issuers/versions.geul",
	} {
		if !isDocPath(p) {
			t.Fatal(p)
		}
	}
	for _, p := range []string{"", "0x00000001/hashes/0xabcd.geul", "0x00000000/hashes/0xab.geul", "0x00000001/docs/a.pub"} {
		if isDocPath(p) {
			t.Fatal(p)
		}
	}
}
