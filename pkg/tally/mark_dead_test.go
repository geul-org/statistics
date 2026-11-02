//ff:func feature=live type=tracker control=sequence
//ff:what TestMarkDead — 같은 경로의 마지막만 살고, 뒤섞인 표도 재생 순서로 가린다
package tally

import "testing"

func TestMarkDead(t *testing.T) {
	d := newDeadSet([]uint64{4})
	marks := []mark{markOf("a", seqOf(0, 2)), markOf("b", seqOf(0, 1)), markOf("a", seqOf(0, 0)), markOf("a", seqOf(0, 3))}
	markDead(marks, d)
	if !d.has(0, 0) || !d.has(0, 2) || d.has(0, 1) || d.has(0, 3) {
		t.Fatalf("got %v", d)
	}
}
