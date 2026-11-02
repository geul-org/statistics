//ff:func feature=count type=counter control=sequence
//ff:what countRecord — 레코드 하나를 세고, 규격이 있으면 inspect.Layout으로 내장 레코드까지 내려간다
package tally

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// countRecord counts r, whose encoding is ws, then the records embedded in
// it. Which payload words are records is known only from the codec's spec:
// data words can look like TAG WORDs, so a payload without a spec is not
// opened and is reported in the returned count.
func countRecord(r core.Record, ws []core.Word, spec codecspec.Spec, c *Counts) uint64 {
	c[slotOf(r.Kind, r.Codec)]++
	if len(r.Payload) == 0 {
		return 0
	}
	fields, ok := spec.Of(r)
	if !ok {
		return 1
	}
	return countParts(inspect.Layout(inspect.Entry{Record: r, Words: ws, Spec: fields}), spec, c)
}
