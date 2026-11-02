//ff:func feature=live type=tracker control=sequence
//ff:what TestDeadSetMark — 해당 비트만 선다
package tally

import "testing"

func TestDeadSetMark(t *testing.T) {
	d := newDeadSet([]uint64{70, 70})
	d.mark(seqOf(1, 65))
	if d[1][1] != 2 || d[0][0] != 0 || d[0][1] != 0 || d[1][0] != 0 {
		t.Fatalf("got %v", d)
	}
}
