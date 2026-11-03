//ff:func feature=archive type=helper control=selection
//ff:what issuerOf — 문서의 Issuer: 공표 주소의 첫 칸, issuers/ 목록은 geul.org 0x00000000, 그 밖에 주소에 없으면 문서를 담은 set의 Issuer
package tally

import "strings"

// issuersDir is the folder of the Issuer list (`issuers/`, §3.13), which
// belongs to geul.org.
const issuersDir = "issuers"

// geulIssuer is geul.org's Issuer ID, the Issuer of the Issuer list (§1.1 96).
const geulIssuer = "0x00000000"

// issuerOf reads the Issuer from a document's address (isDocPath), which
// every kind of document has, version lists without a DOC-ID too. The Issuer
// list `issuers/` is geul.org's. A path without an Issuer, which isDocPath
// does not pass, falls back to the Issuer of the set holding it.
func issuerOf(path, set string) string {
	first, _, nested := strings.Cut(path, "/")
	switch {
	case nested && first == issuersDir:
		return geulIssuer
	case nested && isIssuer(first):
		return first
	default:
		return set
	}
}
