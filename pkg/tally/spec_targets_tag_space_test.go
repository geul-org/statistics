//ff:func feature=archive type=resolver control=iteration dimension=1
//ff:what TestSpecTargetsTagSpace — TAG 공간 512 코드 전수: §3 표의 자리 표지·비 Identified 양화·ISSUER만 빠지고 나머지는 L=0 대상 WORD로 한 번씩
package tally

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSpecTargetsTagSpace(t *testing.T) {
	ts := specTargets()
	for code := 0; code < 0x200; code++ {
		c := code & 0x3F
		var out bool
		switch {
		case code < 0x040: // Reference: markers C<3, 5Q non-Identified, REFERENCE-ISSUER 0x03DL~0x03FL
			out = c < 3 || c >= 5 && c < 0x3C && c%5 != 0 || c >= 0x3D
		case code < 0x080: // Value: VALUE-ISSUER 0x07DL~0x07FL
			out = c >= 0x3D
		case code < 0x1C0: // Edge: EDGE-ISSUER 0x1BDL~0x1BFL
			out = code >= 0x1BD
		}
		w := core.Word(code) << 4
		if n := len(slices.DeleteFunc(slices.Clone(ts), func(x core.Word) bool { return x != w })); out && n != 0 || !out && n != 1 {
			t.Fatalf("0x%03XL: %d times, left out %v", code, n, out)
		}
	}
}
