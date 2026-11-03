//ff:func feature=archive type=resolver control=iteration dimension=1
//ff:what TestSpecTargetsTagSpace — TAG 공간 512 코드 전수: §3.15 표의 자리 표지·비 Identified 양화·ISSUER만 빠지고 나머지는 L=0 대상 WORD로 한 번씩
package tally

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSpecTargetsTagSpace(t *testing.T) {
	ts := specTargets()
	for code := 0; code < 0x200; code++ {
		o := code & 0x1F // offset C in the 32-code section KKKK = code>>5 (§1.1 109)
		var out bool
		switch code >> 5 {
		case 0: // Reference entities: markers C<3, 5Q non-Identified at 5 <= C < 30
			out = o < 3 || o >= 5 && o < 30 && o%5 != 0
		case 1: // Reference EXTERNAL: 5Q non-Identified at C < 30
			out = o < 30 && o%5 != 0
		case 2: // Reference others: REFERENCE-ISSUER 0x05DL~0x05FL
			out = o >= 0x1D
		case 3, 4: // Value: VALUE-ISSUER 0x09DL~0x09FL
			out = code >= 0x09D
		case 5, 6, 7, 8, 9, 10, 11, 12, 13: // Edge: EDGE-ISSUER 0x1BDL~0x1BFL
			out = code >= 0x1BD
		}
		w := core.Word(code) << 4
		if n := len(slices.DeleteFunc(slices.Clone(ts), func(x core.Word) bool { return x != w })); out && n != 0 || !out && n != 1 {
			t.Fatalf("0x%03XL: %d times, left out %v", code, n, out)
		}
	}
}
