//ff:func feature=report type=helper control=sequence
//ff:what TestIssuerCompare — 폭이 먼저, 같은 폭은 값
package tally

import "testing"

func TestIssuerCompare(t *testing.T) {
	if issuerCompare("0xFFFFFFFF", "0x0000000100000000") >= 0 || issuerCompare("0x00000002", "0x00000001") <= 0 || issuerCompare("0x00000001", "0x00000001") != 0 {
		t.Fatal("order")
	}
}
