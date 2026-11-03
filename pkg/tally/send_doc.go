//ff:func feature=count type=helper control=sequence
//ff:what sendDoc — 인벤토리 단계를 걸러 문서 일감을 보낸다, 작업자가 멈췄으면 오류로 돌아온다
package tally

import (
	"context"

	"github.com/geul-org/gogeul/pkg/pack"
)

// sendDoc passes e to the workers unless it is dead, a deletion or no document
// address. A hashes/-only record (`<Issuer>/hashes/`) is only tallied in r: it has no address of its
// own, and it may be a superseded version.
func sendDoc(ctx context.Context, e pack.InventoryEntry, dead bool, jobs chan<- pack.InventoryEntry, r *Report) error {
	if e.Path == "" {
		r.HashOnly++
		return nil
	}
	if dead || e.Deleted || !isDocPath(e.Path) {
		return nil
	}
	select {
	case jobs <- e:
		return nil
	case <-ctx.Done():
		return context.Cause(ctx)
	}
}
