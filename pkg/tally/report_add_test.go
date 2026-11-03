//ff:func feature=report type=helper control=sequence
//ff:what TestReportAdd — Issuer 표와 건수를 더한다
package tally

import "testing"

func TestReportAdd(t *testing.T) {
	a, b := newReport(), newReport()
	a.counts("0x00000001")[1] = 2
	a.Unopened, a.HashOnly = 1, 2
	b.counts("0x00000001")[1] = 3
	b.counts("0x00000002")[4] = 5
	b.Unopened, b.HashOnly = 10, 20
	a.add(b)
	if a.Issuers["0x00000001"][1] != 5 || a.Issuers["0x00000002"][4] != 5 || a.Unopened != 11 || a.HashOnly != 22 {
		t.Fatalf("got %+v", a)
	}
}
