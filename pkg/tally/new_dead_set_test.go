//ff:func feature=live type=tracker control=sequence
//ff:what TestNewDeadSet — 단계 수에 맞는 WORD 수
package tally

import "testing"

func TestNewDeadSet(t *testing.T) {
	d := newDeadSet([]uint64{0, 1, 64, 65})
	if len(d) != 4 || len(d[0]) != 0 || len(d[1]) != 1 || len(d[2]) != 1 || len(d[3]) != 2 {
		t.Fatalf("got %v", d)
	}
}
