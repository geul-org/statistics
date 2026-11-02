//ff:func feature=archive type=helper control=sequence
//ff:what codecDocOf — 시험용 규격 문서: 대상 하나에 레코드 필드 없이 나머지 날값 필드 하나
package tally

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// codecDocOf builds the spec document of target with one rest field.
func codecDocOf(t *testing.T, target core.Word) []byte {
	t.Helper()
	id, err := core.EncodeDocID([]core.Word{0}, []core.Word{0, 0, 0, target})
	if err != nil {
		t.Fatal(err)
	}
	in := wordsOf(t, core.EncodeSpecVersion(core.SpecVersion), id)
	in = append(in, 0x180E)
	in = append(in, wordsOf(t, core.EncodeDeclare(core.Declare{EdgeID: 1, Target: target}),
		core.EncodeFieldDecl(core.FieldDecl{EdgeID: 2, Declare: 1, Width: core.FieldRest}))...)
	return core.PutWords(core.WrapDocument(append(in, 0x180F)))
}
