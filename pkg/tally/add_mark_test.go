//ff:func feature=live type=tracker control=sequence
//ff:what TestAddMark — 문서 경로가 아니면 건너뛰고, 칸이 다르면 건너뛰고, 표가 차면 오류
package tally

import (
	"errors"
	"testing"
)

func TestAddMark(t *testing.T) {
	marks := make([]mark, 1)
	used := 0
	if err := addMark("0x00000001/hashes/0xab.geul", 1, 0, 1, marks, &used); err != nil || used != 0 {
		t.Fatalf("%v %d", err, used)
	}
	other := (markOf("a.geul", 0).h0%2 + 1) % 2
	if err := addMark("a.geul", 2, other, 2, marks, &used); err != nil || used != 0 {
		t.Fatalf("%v %d", err, used)
	}
	if err := addMark("a.geul", 3, 0, 1, marks, &used); err != nil || used != 1 || marks[0].seq != 3 {
		t.Fatalf("%v %d %+v", err, used, marks)
	}
	if err := addMark("b.geul", 4, 0, 1, marks, &used); !errors.Is(err, errMarksFull) {
		t.Fatalf("got %v", err)
	}
}
