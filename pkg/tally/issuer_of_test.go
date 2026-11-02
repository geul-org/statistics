//ff:func feature=archive type=helper control=iteration dimension=1
//ff:what TestIssuerOf — 주소의 첫 칸, 주소에 없으면 set의 Issuer
package tally

import "testing"

func TestIssuerOf(t *testing.T) {
	cases := map[string]string{
		"0x00010000/0x5000000000000006/versions.geul":   "0x00010000",
		"0x00010000/0x5000000000000006/0x00000001.geul": "0x00010000",
		"0x0001/documents.geul":                         "0x0001",
		"issuers.geul":                                  "0x0000",
		"other/name.geul":                               "0x0000",
	}
	for p, want := range cases {
		if got := issuerOf(p, "0x0000"); got != want {
			t.Fatalf("%s: %s", p, got)
		}
	}
}
