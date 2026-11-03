//ff:func feature=archive type=resolver control=iteration dimension=1
//ff:what TestSpecTargets — 직접 코드 456개(Reference 46: 묶음 머리 11·포인터 3·빈칸 32, Value 61, Edge 285, Meta 64), 자리 표지·양화·ISSUER 코드 없음
package tally

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSpecTargets(t *testing.T) {
	ts := specTargets()
	// §3.15: Reference 96 - 3 markers - 44 non-Identified 5Q - 3 ISSUER = 46,
	// Value 64 - 3, Edge 288 - 3, Meta 64.
	if len(ts) != 46+61+285+64 {
		t.Fatalf("got %d", len(ts))
	}
	for _, w := range []core.Word{0x0400, 0x0410, 0x0050, 0x0140, 0x0390, 0x0420, 0x0600, 0x07A0, 0x09C0, 0x0A00, 0x0AA0, 0x1200, 0x1BC0, 0x1C00, 0x1FF0, 0x0030, 0x01E0, 0x03E0, 0x0430} {
		if !slices.Contains(ts, w) {
			t.Fatalf("missing 0x%04X", w)
		}
	}
	for _, w := range []core.Word{0x0000, 0x0010, 0x0020, 0x0060, 0x0150, 0x03D0, 0x05D0, 0x05F0, 0x09D0, 0x09F0, 0x1BD0, 0x1BF0, 0x0210, 0x03A0} {
		if slices.Contains(ts, w) {
			t.Fatalf("0x%04X: want left out", w)
		}
	}
}
