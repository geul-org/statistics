//ff:func feature=count type=helper control=iteration dimension=1
//ff:what TestSlotOfTagSpace — TAG 공간 512 코드 전수: 칸은 코드 그대로이고 서로 다르며, slotKey는 KKK의 Kind와 0x000L로 되돌린다
package tally

import (
	"fmt"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSlotOfTagSpace(t *testing.T) {
	// §2.1 arithmetic, not core constants: KKK = code>>6 gives the Kind and the
	// start of its zone; the Record Codec is the code less that start.
	kinds := [8]core.Kind{core.KindReference, core.KindValue, core.KindEdge, core.KindEdge, core.KindEdge, core.KindEdge, core.KindEdge, core.KindMeta}
	starts := [8]int{0x000, 0x040, 0x080, 0x080, 0x080, 0x080, 0x080, 0x1C0}
	names := [8]string{"REFERENCE", "VALUE", "EDGE", "EDGE", "EDGE", "EDGE", "EDGE", "META"}
	seen := map[int]bool{}
	for code := 0; code < 0x200; code++ {
		k := code >> 6
		slot := slotOf(kinds[k], uint16(code-starts[k]))
		if slot != code || seen[slot] || slot >= len(Counts{}) {
			t.Fatalf("0x%03XL: slot 0x%03X (seen %v)", code, slot, seen[slot])
		}
		seen[slot] = true
		if kind, c := slotKey(slot); kind != names[k] || c != fmt.Sprintf("0x%03XL", code) {
			t.Fatalf("0x%03XL: %s %s", code, kind, c)
		}
	}
}
