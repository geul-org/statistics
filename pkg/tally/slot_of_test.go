//ff:func feature=count type=helper control=sequence
//ff:what TestSlotOf — 칸은 TAG 9비트 코드다: Reference의 양화·자리 표지도 따로, Edge·EDGE-ISSUER와 Meta 끝까지 512칸 안
package tally

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSlotOf(t *testing.T) {
	// TAG IDs from migration1 §3: QID Identified 0x014L, Generic 0x015L; NO-VALUE 0x002L.
	if a, b := slotOf(core.KindReference, core.CodecWikidataQID), slotOf(core.KindReference, core.CodecWikidataQID+1); a != 0x014 || b != 0x015 {
		t.Fatalf("Reference keeps its Quantifier: 0x%03X 0x%03X", a, b)
	}
	if got := slotOf(core.KindReference, 2); got != 0x002 {
		t.Fatalf("slot marker NO-VALUE: 0x%03X", got)
	}
	if a, b := slotOf(core.KindValue, core.CodecUTF8TextEven), slotOf(core.KindValue, core.CodecUTF8TextEven+1); a != 0x072 || b != 0x073 {
		t.Fatalf("UTF8-TEXT even/odd: 0x%03X 0x%03X", a, b)
	}
	if got := slotOf(core.KindEdge, core.EdgeTriple); got != 0x0AA {
		t.Fatalf("TRIPLE: 0x%03X", got)
	}
	if got := slotOf(core.KindEdge, core.CodecEdgeIssuer16Word); got != 0x1BF {
		t.Fatalf("EDGE-ISSUER-16WORD: 0x%03X", got)
	}
	if got := slotOf(core.KindMeta, 0x3F); got != 0x1FF || len(Counts{}) != 0x200 {
		t.Fatalf("last slot: 0x%03X", got)
	}
}
