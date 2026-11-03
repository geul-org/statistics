//ff:func feature=live type=helper control=sequence
//ff:what isDocPath — 문서 주소인가: <Issuer>/hashes/ 조회 주소와 .geul이 아닌 경로는 문서 주소가 아니다
package tally

import "strings"

// isDocPath reports whether path, relative to `.well-known/geul/` (§3.13 공표
// 주소), addresses a document: a version or version list of a document
// (`<Issuer>/docs/<문서 ID>/`) or codec (`<Issuer>/codecs/<코덱 ID>/`), or a
// version or version list of a list (`issuers/`, `<Issuer>/docs/`,
// `<Issuer>/codecs/`). `<Issuer>/hashes/` is a lookup by content hash, also of
// superseded versions, not a document's own address.
func isDocPath(path string) bool {
	_, rest, _ := strings.Cut(path, "/")
	return strings.HasSuffix(path, ".geul") && !strings.HasPrefix(rest, "hashes/")
}
