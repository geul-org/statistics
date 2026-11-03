//ff:func feature=archive type=helper control=iteration dimension=1
//ff:what TestIsIssuer — 폭 8·16·64자리(2·4·16 WORD) A-F 대문자(저장 표기)만 Issuer다, 1 WORD 4자리는 아니다
package tally

import (
	"strings"
	"testing"
)

func TestIsIssuer(t *testing.T) {
	for _, s := range []string{"0x00000000", "0x00000001", "0x00010000", "0x0000000100000000", "0x" + strings.Repeat("AB", 32), "0xABCDABCD"} {
		if !isIssuer(s) {
			t.Fatal(s)
		}
	}
	for _, s := range []string{"", "h", "0x0000", "0x0001", "0x000", "0x000000", "0X00000000", "0xabcdabcd", "0x" + strings.Repeat("ab", 32), "geul.org", "issuers", "0x" + strings.Repeat("A", 32)} {
		if isIssuer(s) {
			t.Fatal(s)
		}
	}
}
