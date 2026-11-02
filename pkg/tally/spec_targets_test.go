//ff:func feature=archive type=resolver control=sequence
//ff:what TestSpecTargets — 직접 코드 전부, 자리 표지와 ISSUER 코드 없음
package tally

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSpecTargets(t *testing.T) {
	ts := specTargets()
	if len(ts) != 15+3*0x7C {
		t.Fatalf("got %d", len(ts))
	}
	if !slices.Contains(ts, core.TargetWord(core.KindEdge, core.EdgeTriple)) || !slices.Contains(ts, core.TargetWord(core.KindReference, 0x78)) {
		t.Fatal("missing target")
	}
	if slices.Contains(ts, core.TargetWord(core.KindReference, 0)) || slices.Contains(ts, core.TargetWord(core.KindValue, 0x7C)) {
		t.Fatal("reserved target")
	}
}
