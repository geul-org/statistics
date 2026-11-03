//ff:func feature=count type=helper control=sequence
//ff:what testSpec — 시험용 규격: TRIPLE과 LOCAL·QID·PID
package tally

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
)

// testSpec declares TRIPLE (three record fields, the Object also taking the
// slot markers, and a raw rank), DOC-1WORD and the References the fixtures use
// (a raw rest field each).
func testSpec() codecspec.Spec {
	id := []codecspec.Field{{Name: "id", Width: int(core.FieldRest)}}
	qid := core.TargetWord(core.KindReference, core.CodecWikidataQID)
	return codecspec.Spec{
		core.TargetWord(core.KindEdge, core.EdgeTriple): {
			{Name: "predicate"}, {Name: "subject"}, {Name: "object", Codecs: []core.Word{core.Absent, core.UnknownValue, core.NoValue, qid}}, {Name: "rank", Width: 1},
		},
		core.TargetWord(core.KindMeta, core.MetaDoc1Word):    id,
		core.TargetWord(core.KindReference, core.CodecLocal): id,
		qid: id,
		core.TargetWord(core.KindReference, core.CodecWikidataPID): id,
	}
}
