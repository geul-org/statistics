//ff:func feature=archive type=helper control=iteration dimension=1
//ff:what TestIssuerOf — 주소의 첫 칸(2·4·16 WORD), issuers/ 목록은 geul.org 0x00000000, 주소에 없으면 set의 Issuer
package tally

import "testing"

func TestIssuerOf(t *testing.T) {
	for p, want := range map[string]string{
		"0x00000001/docs/0x5000000000000006/versions.geul":   "0x00000001",
		"0x00000001/docs/0x5000000000000006/0x00000001.geul": "0x00000001",
		"0x00000002/docs/0x00000001.geul":                    "0x00000002",
		"0x0000000100000000/codecs/0x00000001.geul":          "0x0000000100000000",
		"issuers/0x00000001.geul":                            "0x00000000",
		"0x0001/docs/0x00000001.geul":                        "0x00000009",
		"other/name.geul":                                    "0x00000009",
		"name.geul":                                          "0x00000009",
	} {
		if got := issuerOf(p, "0x00000009"); got != want {
			t.Fatalf("%s: %s want %s", p, got, want)
		}
	}
}
