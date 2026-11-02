//ff:type feature=archive type=model
//ff:what Source — pack set 하나에서 세기에 쓰는 읽기, *pack.Reader가 구현한다
package tally

import (
	"context"

	"github.com/geul-org/gogeul/pkg/pack"
)

// Source is the part of *pack.Reader counting uses.
type Source interface {
	DataPacks() []uint32
	PackInventory(ctx context.Context, obj uint32, fn func(pack.InventoryEntry) error) error
	Record(ctx context.Context, l pack.Loc) ([]byte, error)
}
