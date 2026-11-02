//ff:func feature=live type=tracker control=sequence
//ff:what TestMarkCompare — h0, h1, seq 순으로 비교한다
package tally

import "testing"

func TestMarkCompare(t *testing.T) {
	if markCompare(mark{h0: 1, h1: 9, seq: 9}, mark{h0: 2}) >= 0 {
		t.Fatal("h0")
	}
	if markCompare(mark{h0: 1, h1: 1, seq: 9}, mark{h0: 1, h1: 2}) >= 0 {
		t.Fatal("h1")
	}
	if markCompare(mark{h0: 1, h1: 1, seq: 2}, mark{h0: 1, h1: 1, seq: 1}) <= 0 {
		t.Fatal("seq")
	}
}
