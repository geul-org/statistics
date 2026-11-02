//ff:func feature=archive type=helper control=sequence
//ff:what TestSetIssuer — Issuer 폴더면 그 이름, 아니면 없음 표시
package tally

import "testing"

func TestSetIssuer(t *testing.T) {
	if got := setIssuer("/a/.well-known/geul/0x0000"); got != "0x0000" {
		t.Fatal(got)
	}
	if got := setIssuer("/a/.well-known/geul"); got != noIssuer {
		t.Fatal(got)
	}
}
