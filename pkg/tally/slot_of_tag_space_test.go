//ff:func feature=count type=helper control=iteration dimension=1
//ff:what TestSlotOfTagSpace — TAG 공간 512 코드 전수: 칸은 코드 그대로이고 서로 다르며, slotKey는 섹션 KKKK의 Kind와 0x000L로 되돌린다
package tally

import (
	"fmt"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSlotOfTagSpace(t *testing.T) {
	// migration1 §3.15 section table, not core constants: the section KKKK =
	// code>>5 gives the Kind (Reference 0~2, Value 3·4, Edge 5~13, Meta 14·15)
	// and the start of its zone; the Record Codec is the code less that start.
	const r, v, e, m = core.KindReference, core.KindValue, core.KindEdge, core.KindMeta
	kinds := [16]core.Kind{r, r, r, v, v, e, e, e, e, e, e, e, e, e, m, m}
	starts := map[core.Kind]int{r: 0x000, v: 0x060, e: 0x0A0, m: 0x1C0}
	names := map[core.Kind]string{r: "REFERENCE", v: "VALUE", e: "EDGE", m: "META"}
	seen := map[int]bool{}
	for code := 0; code < 0x200; code++ {
		k := kinds[code>>5]
		slot := slotOf(k, uint16(code-starts[k]))
		if slot != code || seen[slot] || slot >= len(Counts{}) {
			t.Fatalf("0x%03XL: slot 0x%03X (seen %v)", code, slot, seen[slot])
		}
		seen[slot] = true
		if kind, c := slotKey(slot); kind != names[k] || c != fmt.Sprintf("0x%03XL", code) {
			t.Fatalf("0x%03XL: %s %s", code, kind, c)
		}
	}
}
