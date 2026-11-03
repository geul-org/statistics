//ff:func feature=report type=helper control=sequence
//ff:what TestReportCounts — 없으면 만들고 있으면 같은 표를 준다
package tally

import "testing"

func TestReportCounts(t *testing.T) {
	r := newReport()
	a := r.counts("0x00000001")
	a[3] = 7
	if b := r.counts("0x00000001"); b != a || b[3] != 7 || len(r.Issuers) != 1 {
		t.Fatalf("got %v", r.Issuers)
	}
}
