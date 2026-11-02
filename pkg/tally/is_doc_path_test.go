//ff:func feature=live type=helper control=iteration dimension=1
//ff:what TestIsDocPath — 문서 주소와 아닌 것
package tally

import "testing"

func TestIsDocPath(t *testing.T) {
	for _, p := range []string{"0x0001/0x0000000000000001/versions.geul", "0x0001/documents.geul", "issuers.geul"} {
		if !isDocPath(p) {
			t.Fatal(p)
		}
	}
	for _, p := range []string{"", "h/0xabcd.geul", "keys/a.pub"} {
		if isDocPath(p) {
			t.Fatal(p)
		}
	}
}
