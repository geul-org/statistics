//ff:func feature=live type=tracker control=sequence
//ff:what TestDeadSetHas — 표시한 단계만 죽었다
package tally

import "testing"

func TestDeadSetHas(t *testing.T) {
	d := newDeadSet([]uint64{70})
	d.mark(seqOf(0, 3))
	if !d.has(0, 3) || d.has(0, 4) || d.has(0, 67) {
		t.Fatalf("got %v", d)
	}
}
