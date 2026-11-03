//ff:func feature=archive type=resolver control=iteration dimension=1
//ff:what TestSpecTargets — 직접 코드 456개(포인터·묶음 머리·CODEC-POINTER 14, Value 61, Edge 317, Meta 64), 자리 표지·양화·ISSUER 코드 없음
package tally

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSpecTargets(t *testing.T) {
	ts := specTargets()
	// §3: Reference 64 - 3 markers - 44 non-Identified 5Q - 3 ISSUER = 14,
	// Value 64 - 3, Edge 320 - 3, Meta 64.
	if len(ts) != 14+61+317+64 {
		t.Fatalf("got %d", len(ts))
	}
	for _, w := range []core.Word{0x0030, 0x0040, 0x0050, 0x0140, 0x0370, 0x03C0, 0x0400, 0x05A0, 0x07C0, 0x0800, 0x08A0, 0x1000, 0x1BC0, 0x1C00, 0x1FF0} {
		if !slices.Contains(ts, w) {
			t.Fatalf("missing 0x%04X", w)
		}
	}
	for _, w := range []core.Word{0x0000, 0x0010, 0x0020, 0x0060, 0x0150, 0x03B0, 0x03D0, 0x03F0, 0x07D0, 0x07F0, 0x1BD0, 0x1BF0} {
		if slices.Contains(ts, w) {
			t.Fatalf("0x%04X: want left out", w)
		}
	}
}
