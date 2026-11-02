//ff:func feature=count type=helper control=sequence
//ff:what tripleOf — 시험용 TRIPLE: Edge Id, Predicate, Subject, Object, Rank 날값
package tally

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// tripleOf builds a TRIPLE with Edge Id id and rank 0.
func tripleOf(t *testing.T, id uint64, pred, subj, obj core.Record) core.Record {
	t.Helper()
	p := append(wordsOf(t, core.EncodeLocalRef(id), pred, subj, obj), 0)
	return core.Record{Kind: core.KindEdge, Codec: core.EdgeTriple, Payload: p}
}
