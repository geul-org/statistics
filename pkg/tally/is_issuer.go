//ff:func feature=archive type=helper control=sequence
//ff:what isIssuer — 정규형 Issuer ID인가: 0x와 소문자 hex 4·8·16·64자리
package tally

import "regexp"

// issuerID is publishing.md's Issuer notation.
var issuerID = regexp.MustCompile(`^0x([0-9a-f]{4}|[0-9a-f]{8}|[0-9a-f]{16}|[0-9a-f]{64})$`)

// isIssuer reports whether s is a canonical Issuer ID.
func isIssuer(s string) bool {
	return issuerID.MatchString(s)
}
