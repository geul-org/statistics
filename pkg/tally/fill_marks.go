//ff:func feature=live type=tracker control=iteration dimension=1
//ff:what fillMarks — 모든 인벤토리를 pack 차례로 읽어 칸 part에 든 경로의 줄을 표에 채운다
package tally

import (
	"context"
	"errors"

	"github.com/geul-org/gogeul/pkg/pack"
)

// errMarksFull means one path repeats so often that a partition overflowed.
var errMarksFull = errors.New("path tracking table is full: raise -memory")

// fillMarks fills marks with the steps of partition part of parts and returns
// how many.
func fillMarks(ctx context.Context, src Source, packs []uint32, part, parts uint64, marks []mark) (int, error) {
	used := 0
	for i, obj := range packs {
		step := uint32(0)
		err := src.PackInventory(ctx, obj, func(e pack.InventoryEntry) error {
			step++
			return addMark(e.Path, seqOf(i, step-1), part, parts, marks, &used)
		})
		if err != nil {
			return 0, err
		}
	}
	return used, nil
}
