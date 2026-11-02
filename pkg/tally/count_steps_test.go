//ff:func feature=live type=tracker control=sequence
//ff:what TestCountSteps — pack별 단계 수와 문서 경로 수, 인벤토리 오류
package tally

import (
	"context"
	"errors"
	"slices"
	"testing"
)

func TestCountSteps(t *testing.T) {
	s := twoPackSource()
	steps, tracked, err := countSteps(context.Background(), s, s.packs)
	if err != nil || !slices.Equal(steps, []uint64{5, 7}) || tracked != 9 {
		t.Fatalf("%v %d %v", steps, tracked, err)
	}
	s.invCall, s.invFail, s.err = new(int), 2, errors.New("broken")
	if _, _, err := countSteps(context.Background(), s, s.packs); !errors.Is(err, s.err) {
		t.Fatalf("got %v", err)
	}
}
