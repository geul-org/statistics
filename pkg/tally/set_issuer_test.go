//ff:func feature=archive type=helper control=sequence
//ff:what TestSetIssuer — Issuer 폴더(2·4·16 WORD)면 그 이름, 옛 1 WORD 이름이나 다른 폴더면 없음 표시
package tally

import "testing"

func TestSetIssuer(t *testing.T) {
	if got := setIssuer("/a/.well-known/geul/0x00000000"); got != "0x00000000" {
		t.Fatal(got)
	}
	if got := setIssuer("/a/.well-known/geul/0x0000"); got != noIssuer {
		t.Fatal(got)
	}
	if got := setIssuer("/a/.well-known/geul"); got != noIssuer {
		t.Fatal(got)
	}
}
