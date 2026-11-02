//ff:func feature=archive type=helper control=iteration dimension=1
//ff:what TestIsIssuer — 폭 4·8·16·64자리 소문자만 Issuer다
package tally

import (
	"strings"
	"testing"
)

func TestIsIssuer(t *testing.T) {
	for _, s := range []string{"0x0000", "0x00010000", "0x0000000100000000", "0x" + strings.Repeat("ab", 32)} {
		if !isIssuer(s) {
			t.Fatal(s)
		}
	}
	for _, s := range []string{"", "h", "0x000", "0x000000", "0X0000", "0xABCD", "geul.org", "issuers.geul"} {
		if isIssuer(s) {
			t.Fatal(s)
		}
	}
}
