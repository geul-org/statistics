//ff:func feature=report type=helper control=sequence
//ff:what TestCountsAdd — 칸별로 더한다
package tally

import "testing"

func TestCountsAdd(t *testing.T) {
	a, b := Counts{0: 1, 511: 2}, Counts{0: 10, 7: 3}
	a.add(&b)
	if a[0] != 11 || a[7] != 3 || a[511] != 2 {
		t.Fatalf("got %v", nonZero(a))
	}
}
