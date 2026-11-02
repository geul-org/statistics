//ff:type feature=archive type=model
//ff:what fakeSource — 시험용 Source: 메모리의 인벤토리와 문서, 정해진 오류
package tally

import "github.com/geul-org/gogeul/pkg/pack"

// fakeSource serves steps per object and docs per Loc.Off. invFail is the
// PackInventory call (1-based) that fails with err; recFail fails Record.
type fakeSource struct {
	packs   []uint32
	steps   map[uint32][]pack.InventoryEntry
	docs    map[uint64][]byte
	invFail int
	invCall *int
	recFail bool
	err     error
}
