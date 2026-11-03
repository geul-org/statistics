//ff:func feature=live type=helper control=iteration dimension=1
//ff:what TestIsDocPath — 문서·코덱·Issuer 목록의 버전·versions.geul·latest.geul은 문서 주소, <Issuer>/hashes/·옛 docs·codecs 목록·형식 밖 경로는 아니다
package tally

import (
	"strings"
	"testing"
)

func TestIsDocPath(t *testing.T) {
	i64 := "0x" + strings.Repeat("0", 63) + "1"
	for _, p := range []string{
		"0x00000001/docs/0x0000000000000001/versions.geul",
		"0x00000001/docs/0x0000000000000001/0x00000001.geul",
		"0x00000001/docs/0x0000000000000001/latest.geul",
		"0x00000000/codecs/0x00000140/0x00000001.geul",
		"0x00000000/codecs/0x00000140/versions.geul",
		"0x0000000100000000/docs/0x0000000000000001/0x00000002.geul",
		i64 + "/codecs/0x0000c000/0x00000001.geul",
		"issuers/0x00000001.geul",
		"issuers/versions.geul",
		"issuers/latest.geul",
	} {
		if !isDocPath(p) {
			t.Fatal(p)
		}
	}
	for _, p := range []string{
		"", "0x00000001/hashes/0xabcd.geul", "0x00000000/hashes/0xab.geul", "0x00000001/docs/a.pub",
		"0x00000001/docs/0x00000001.geul", "0x00000001/docs/versions.geul", "0x00000001/docs/latest.geul",
		"0x00000000/codecs/0x00000001.geul", "0x00000000/codecs/versions.geul", "0x00000000/codecs/latest.geul",
		"0x00000001/docs/0x0000000000000001/0x00000000.geul", "0x00000001/docs/0x00000001/versions.geul",
		"0x00000000/codecs/0x0000000000000140/versions.geul", "0x0001/docs/0x0000000000000001/versions.geul",
		"0x00000001/docs/0x000000000000000A/versions.geul", "issuers/x/versions.geul", "issuers.geul",
		"0x0000/documents.geul", "0x00000001/keys/0x00000001/versions.geul",
	} {
		if isDocPath(p) {
			t.Fatal(p)
		}
	}
}
