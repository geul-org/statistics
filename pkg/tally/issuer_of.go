//ff:func feature=archive type=helper control=sequence
//ff:what issuerOf — 문서의 Issuer: 공표 주소의 첫 칸, 주소에 없으면 문서를 담은 set의 Issuer
package tally

import "strings"

// issuerOf reads the Issuer from a document's address, which every kind of
// document has, version lists without a DOC-ID too. issuers.geul has none in
// its address; it then belongs to the Issuer of the set holding it.
func issuerOf(path, set string) string {
	first, _, nested := strings.Cut(path, "/")
	if nested && isIssuer(first) {
		return first
	}
	return set
}
