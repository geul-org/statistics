//ff:func feature=live type=tracker control=iteration dimension=1
//ff:what TestFillMarks — 칸 합이 문서 경로 수와 같고, 재생 순번을 쓰고, 표 넘침과 인벤토리 오류는 돌아온다
package tally

import (
	"context"
	"errors"
	"testing"
)

func TestFillMarks(t *testing.T) {
	ctx := context.Background()
	s := twoPackSource()
	marks := make([]mark, 9)
	total := 0
	for part := range uint64(3) {
		n, err := fillMarks(ctx, s, s.packs, part, 3, marks)
		if err != nil {
			t.Fatal(err)
		}
		total += n
	}
	if total != 9 {
		t.Fatalf("got %d", total)
	}
	n, err := fillMarks(ctx, s, s.packs, 0, 1, marks)
	if err != nil || n != 9 || marks[8].seq != seqOf(1, 3) {
		t.Fatalf("%d %v %+v", n, err, marks[8])
	}
	if _, err := fillMarks(ctx, s, s.packs, 0, 1, marks[:8]); !errors.Is(err, errMarksFull) {
		t.Fatalf("got %v", err)
	}
}
