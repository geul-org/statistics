//ff:func feature=archive type=helper control=iteration dimension=1
//ff:what fakeSource.PackInventory — 시험용 인벤토리 재생, invFail번째 호출은 실패한다
package tally

import (
	"context"

	"github.com/geul-org/gogeul/pkg/pack"
)

// PackInventory replays obj's steps.
func (s fakeSource) PackInventory(ctx context.Context, obj uint32, fn func(pack.InventoryEntry) error) error {
	if s.invCall != nil {
		*s.invCall++
		if *s.invCall == s.invFail {
			return s.err
		}
	}
	for _, e := range s.steps[obj] {
		if err := fn(e); err != nil {
			return err
		}
	}
	return nil
}
