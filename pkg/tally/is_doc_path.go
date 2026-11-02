//ff:func feature=live type=helper control=sequence
//ff:what isDocPath — 문서 주소인가: h/ 조회 주소와 .geul이 아닌 경로는 문서 주소가 아니다
package tally

import "strings"

// isDocPath reports whether path addresses a document. h/ is a lookup by
// content hash, also of superseded versions, not a document's own address.
func isDocPath(path string) bool {
	return strings.HasSuffix(path, ".geul") && !strings.HasPrefix(path, "h/")
}
