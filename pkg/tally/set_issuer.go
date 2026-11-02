//ff:func feature=archive type=helper control=sequence
//ff:what setIssuer — set 폴더 이름이 Issuer면 그 이름, 아니면 Issuer 없음 표시
package tally

import "path/filepath"

// noIssuer stands in the issuer column for a document with no Issuer anywhere.
const noIssuer = "-"

// setIssuer names the Issuer of the set in dir. An Issuer folder given
// directly opens as a top set with an empty name, so the folder name is read.
func setIssuer(dir string) string {
	if name := filepath.Base(dir); isIssuer(name) {
		return name
	}
	return noIssuer
}
