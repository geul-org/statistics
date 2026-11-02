//ff:func feature=count type=helper control=sequence
//ff:what TestSlotOf — Reference는 하위 3비트를 버리고 다른 Kind는 7비트를 그대로 둔다
package tally

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSlotOf(t *testing.T) {
	if a, b := slotOf(core.KindReference, 0x20), slotOf(core.KindReference, 0x27); a != b || a != 4 {
		t.Fatalf("Reference low bits: %d %d", a, b)
	}
	if got := slotOf(core.KindReference, 0x02); got != 0 {
		t.Fatalf("slot marker: %d", got)
	}
	if a, b := slotOf(core.KindValue, 0x40), slotOf(core.KindValue, 0x41); a == b || a != 1<<7|0x40 {
		t.Fatalf("Value keeps 7 bits: %d %d", a, b)
	}
	if got := slotOf(core.KindEdge, 0x7C); got != 2<<7|0x7C {
		t.Fatalf("ISSUER code as is: %d", got)
	}
	if got := slotOf(core.KindMeta, 0x7F); got != len(Counts{})-1 {
		t.Fatalf("last slot: %d", got)
	}
}
