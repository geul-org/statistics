//ff:func feature=report type=helper control=sequence
//ff:what TestSlotKey — Reference는 hex 한 자리, 나머지는 두 자리
package tally

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSlotKey(t *testing.T) {
	if k, c := slotKey(slotOf(core.KindReference, 0x20)); k != "REFERENCE" || c != "0x4" {
		t.Fatalf("got %s %s", k, c)
	}
	if k, c := slotKey(slotOf(core.KindEdge, 0x0A)); k != "EDGE" || c != "0x0A" {
		t.Fatalf("got %s %s", k, c)
	}
	if k, c := slotKey(slotOf(core.KindValue, 0x7C)); k != "VALUE" || c != "0x7C" {
		t.Fatalf("got %s %s", k, c)
	}
}
