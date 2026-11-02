//ff:func feature=live type=tracker control=iteration dimension=1
//ff:what TestFindDead — 한도 없을 때와 칸을 나눌 때 결과가 같다, 두 단계의 인벤토리 오류
package tally

import (
	"context"
	"errors"
	"testing"
)

func TestFindDead(t *testing.T) {
	ctx := context.Background()
	s := twoPackSource()
	for _, memory := range []int64{0, 16 * markSize} {
		d, err := findDead(ctx, s, s.packs, memory)
		if err != nil {
			t.Fatalf("memory %d: %v", memory, err)
		}
		checkDeadSteps(t, d)
	}
	s.invCall, s.invFail, s.err = new(int), 1, errors.New("broken")
	if _, err := findDead(ctx, s, s.packs, 0); !errors.Is(err, s.err) {
		t.Fatalf("first pass: %v", err)
	}
	s.invCall, s.invFail = new(int), 3
	if _, err := findDead(ctx, s, s.packs, 0); !errors.Is(err, s.err) {
		t.Fatalf("second pass: %v", err)
	}
}
