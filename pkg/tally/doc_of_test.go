//ff:func feature=count type=helper control=sequence
//ff:what docOf — 시험용 문서 바이트: DOC-1WORD, HEAD 없이 BODY 레코드들, END
package tally

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// docOf wraps rs into a document without BODY marker, as an empty HEAD allows.
func docOf(t *testing.T, rs ...core.Record) []byte {
	t.Helper()
	return core.PutWords(core.WrapDocument(append(wordsOf(t, rs...), core.TargetWord(core.KindMeta, core.MetaEnd))))
}
