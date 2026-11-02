//ff:func feature=count type=helper control=iteration dimension=1
//ff:what wordsOf — 시험용: 레코드들을 이어 붙인 WORD
package tally

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// wordsOf encodes rs in order.
func wordsOf(t *testing.T, rs ...core.Record) []core.Word {
	t.Helper()
	var out []core.Word
	for _, r := range rs {
		ws, err := core.EncodeRecord(r)
		if err != nil {
			t.Fatal(err)
		}
		out = append(out, ws...)
	}
	return out
}
