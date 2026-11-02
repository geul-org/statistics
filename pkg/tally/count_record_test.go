//ff:func feature=count type=counter control=sequence
//ff:what TestCountRecord — Payload 없음, 규격 없음은 열지 않음, 규격 있음은 내장까지
package tally

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
)

func TestCountRecord(t *testing.T) {
	var c Counts
	marker := refOf(0, 1)
	if u := countRecord(marker, wordsOf(t, marker), testSpec(), &c); u != 0 || c[0] != 1 {
		t.Fatalf("marker: %d %d", u, c[0])
	}
	tr := tripleOf(t, 1, refOf(core.RefWikidataPID, 0, 31), refOf(core.RefWikidataQID, 0, 42), refOf(core.RefWikidataQID, 0, 5))
	c = Counts{}
	if u := countRecord(tr, wordsOf(t, tr), codecspec.Spec{}, &c); u != 1 || len(nonZero(c)) != 1 {
		t.Fatalf("no spec: %d %v", u, nonZero(c))
	}
	c = Counts{}
	if u := countRecord(tr, wordsOf(t, tr), testSpec(), &c); u != 0 || c[slotOf(core.KindReference, core.RefWikidataQID<<3)] != 2 {
		t.Fatalf("spec: %d %v", u, nonZero(c))
	}
}
