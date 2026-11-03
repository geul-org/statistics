//ff:func feature=live type=helper control=sequence
//ff:what twoPackSource — 시험용: 덮어쓰기·삭제·삭제 뒤 재기록·h/ 항목이 든 pack 둘
package tally

import "github.com/geul-org/gogeul/pkg/pack"

// twoPackSource replays: pack 0 sets a, b, c and d, then replaces a in the
// same pack; pack 1 replaces b, deletes c, deletes then sets d again. The h/
// deletion and the h/-only record name no document address.
func twoPackSource() fakeSource {
	l := func(off uint64) pack.Loc { return pack.Loc{Off: off} }
	return fakeSource{packs: []uint32{5, 9}, steps: map[uint32][]pack.InventoryEntry{
		5: {
			{Path: "a.geul", Loc: l(1)}, {Path: "b.geul", Loc: l(2)}, {Path: "c.geul", Loc: l(3)},
			{Path: "d.geul", Loc: l(4)}, {Path: "a.geul", Loc: l(5)},
		},
		9: {
			{Path: "b.geul", Loc: l(6)}, {Path: "c.geul", Deleted: true}, {Path: "d.geul", Deleted: true},
			{Path: "d.geul", Loc: l(7)}, {Path: "0x00000001/hashes/0xab.geul", Deleted: true}, {Loc: l(8), Hashed: true},
			{Path: "keys/a.pub", Loc: l(9)},
		},
	}}
}
