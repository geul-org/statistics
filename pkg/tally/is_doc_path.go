//ff:func feature=live type=helper control=sequence
//ff:what isDocPath — 문서 주소인가: 공표 주소 10가지 중 issuers/와 문서·코덱의 버전·versions.geul·latest.geul; <Issuer>/hashes/와 형식 밖 경로(옛 docs·codecs 목록 포함)는 아니다
package tally

import (
	"regexp"
	"strings"
)

// docPath is the document part of the 10 addresses of §3.13 (`hashes/` aside):
// the Issuer list `issuers/<leaf>`, a document `<Issuer>/docs/<문서 ID 16자리>/<leaf>`
// and a codec `<Issuer>/codecs/<코덱 ID 8자리>/<leaf>`. There is no docs or
// codecs list (§1.1 101).
var docPath = regexp.MustCompile(`^(issuers|0x([0-9A-F]{8}|[0-9A-F]{16}|[0-9A-F]{64})/(docs/0x[0-9A-F]{16}|codecs/0x[0-9A-F]{8}))/(latest|versions|0x[0-9A-F]{8})\.geul$`)

// isDocPath reports whether path, relative to `.well-known/geul/`, addresses a
// document. Version IDs start at 1. `<Issuer>/hashes/` is a lookup by content
// hash, also of superseded versions, not a document's own address.
func isDocPath(path string) bool {
	return docPath.MatchString(path) && !strings.HasSuffix(path, "/0x00000000.geul")
}
