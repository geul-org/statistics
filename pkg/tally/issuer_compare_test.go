//ff:func feature=report type=helper control=sequence
//ff:what TestIssuerCompare — 폭이 먼저, 같은 폭은 값
package tally

import "testing"

func TestIssuerCompare(t *testing.T) {
	if issuerCompare("0xffff", "0x00000000") >= 0 || issuerCompare("0x0002", "0x0001") <= 0 || issuerCompare("0x0001", "0x0001") != 0 {
		t.Fatal("order")
	}
}
