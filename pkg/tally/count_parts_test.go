//ff:func feature=count type=counter control=sequence
//ff:what TestCountParts — 내장 레코드는 세고, 머리와 날값은 넘기고, 규격에 안 맞는 조각과 깨진 조각은 알린다
package tally

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestCountParts(t *testing.T) {
	var c Counts
	parts := []inspect.Part{
		{Name: "TAG", Words: []core.Word{0x0800}, Head: true},
		{Name: "rank", Words: []core.Word{0}},
		{Name: "ID", Words: wordsOf(t, core.EncodeLocalRef(1)), Record: true},
		{Name: genericPayload, Words: []core.Word{0x0051}},
		{Name: "broken", Words: []core.Word{0x2000}, Record: true},
	}
	if u := countParts(parts, testSpec(), &c); u != 2 {
		t.Fatalf("unopened %d", u)
	}
	if n := nonZero(c); len(n) != 1 || c[slotOf(core.KindReference, core.CodecLocal)] != 1 {
		t.Fatalf("counts %v", n)
	}
}
