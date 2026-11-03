//ff:func feature=report type=helper control=iteration dimension=1
//ff:what TestSlotKey — kind는 섹션 KKKK로, codec은 TAG ID 0x000L 표기로
package tally

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSlotKey(t *testing.T) {
	for _, c := range []struct {
		slot       int
		kind, code string
	}{
		{slotOf(core.KindReference, core.CodecWikidataQID), "REFERENCE", "0x014L"},
		{slotOf(core.KindReference, 0), "REFERENCE", "0x000L"},
		{slotOf(core.KindValue, core.CodecLanguageShortCode), "VALUE", "0x07AL"},
		{slotOf(core.KindEdge, core.EdgeTriple), "EDGE", "0x0AAL"},
		{slotOf(core.KindEdge, core.CodecEdgeIssuer2Word), "EDGE", "0x1BDL"},
		{slotOf(core.KindMeta, core.MetaEnd), "META", "0x1C4L"},
	} {
		if k, code := slotKey(c.slot); k != c.kind || code != c.code {
			t.Fatalf("slot %d: %s %s want %s %s", c.slot, k, code, c.kind, c.code)
		}
	}
}
