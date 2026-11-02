//ff:func feature=report type=model control=sequence
//ff:what TestNewReport — 빈 Report는 표를 바로 쓸 수 있다
package tally

import "testing"

func TestNewReport(t *testing.T) {
	r := newReport()
	if r.Issuers == nil || len(r.Issuers) != 0 || r.Unopened != 0 || r.HashOnly != 0 {
		t.Fatalf("got %+v", r)
	}
}
