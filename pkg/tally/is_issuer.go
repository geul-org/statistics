//ff:func feature=archive type=helper control=sequence
//ff:what isIssuer — 정규형 Issuer ID인가: 0x와 A-F 대문자 hex 8·16·64자리(저장 표기, §1.1 120)(2·4·16 WORD)
package tally

import "regexp"

// issuerID is publishing.md's Issuer notation: 2, 4 or 16 words (§3.11, the
// 1-word width is gone).
var issuerID = regexp.MustCompile(`^0x([0-9A-F]{8}|[0-9A-F]{16}|[0-9A-F]{64})$`)

// isIssuer reports whether s is an Issuer ID.
func isIssuer(s string) bool {
	return issuerID.MatchString(s)
}
