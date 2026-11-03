//ff:func feature=count type=counter control=sequence
//ff:what TestCountDocument — 손으로 아는 문서의 합계: Reference 양화·자리 표지는 각자 칸, LLLL은 버림, 구조 레코드, 내장 레코드, 깨진 입력
package tally

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCountDocument(t *testing.T) {
	pred := refOf(core.CodecWikidataPID, 0, 31)
	// QID 1 WORD and 2 WORD target (LLLL differs), Identified and Generic;
	// the objects are the slot markers ABSENT, UNKNOWN-VALUE and NO-VALUE.
	b := docOf(t,
		tripleOf(t, 1, pred, refOf(core.CodecWikidataQID, 0, 42), refOf(0, 0)),
		tripleOf(t, 2, pred, refOf(core.CodecWikidataQID, 1, 1, 0), refOf(0, 1)),
		tripleOf(t, 3, pred, refOf(core.CodecWikidataQID, 0, 42), refOf(0, 2)),
		core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("ab")),
		core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("abcdef")),
	)
	var c Counts
	unopened, err := CountDocument(b, testSpec(), &c)
	if err != nil {
		t.Fatal(err)
	}
	if unopened != 2 { // the two UTF8-TEXT payloads: no spec in testSpec
		t.Fatalf("unopened %d", unopened)
	}
	want := Counts{}
	want[slotOf(core.KindMeta, core.MetaDoc1Word)] = 1
	want[slotOf(core.KindMeta, core.MetaEnd)] = 1
	want[slotOf(core.KindEdge, core.EdgeTriple)] = 3
	want[slotOf(core.KindReference, core.CodecLocal)] = 3
	want[slotOf(core.KindReference, core.CodecWikidataPID)] = 3
	want[slotOf(core.KindReference, core.CodecWikidataQID)] = 2   // 0x014L, 1 WORD and 2 WORD target alike
	want[slotOf(core.KindReference, core.CodecWikidataQID+1)] = 1 // 0x015L Generic: the Quantifier is kept
	want[0x000], want[0x001], want[0x002] = 1, 1, 1               // ABSENT, UNKNOWN-VALUE, NO-VALUE

	want[slotOf(core.KindValue, core.CodecUTF8TextEven)] = 2
	if c != want {
		t.Fatalf("counts differ:\n%v\n%v", nonZero(c), nonZero(want))
	}
	if _, err := CountDocument([]byte{1}, testSpec(), &c); err == nil {
		t.Fatal("odd bytes: want error")
	}
	if _, err := CountDocument(core.PutWords([]core.Word{0x2000}), testSpec(), &c); err == nil {
		t.Fatal("bad prefix: want error")
	}
}
