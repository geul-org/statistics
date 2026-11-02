//ff:func feature=live type=tracker control=iteration dimension=1
//ff:what countSteps — 인벤토리를 한 번 읽어 pack별 단계 수와 문서 경로 단계의 합을 낸다
package tally

import (
	"context"

	"github.com/geul-org/gogeul/pkg/pack"
)

// countSteps sizes the tracking: the steps of each pack in packs and how many
// of all the steps name a document path. Packs are read one at a time, as an
// inventory stays in memory while it is read, up to 300 MB on Wikidata.
func countSteps(ctx context.Context, src Source, packs []uint32) ([]uint64, uint64, error) {
	steps := make([]uint64, len(packs))
	var tracked uint64
	for i, obj := range packs {
		err := src.PackInventory(ctx, obj, func(e pack.InventoryEntry) error {
			steps[i]++
			if isDocPath(e.Path) {
				tracked++
			}
			return nil
		})
		if err != nil {
			return nil, 0, err
		}
	}
	return steps, tracked, nil
}
