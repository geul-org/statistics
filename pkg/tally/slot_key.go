//ff:func feature=report type=helper control=sequence
//ff:what slotKey — 칸 번호를 TSV의 kind와 codec 글자로 되돌린다
package tally

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// slotKey names a slot: the Kind's spec name and the Codec in hex, one digit
// for a Reference's 4 bits and two for the other Kinds' 7 bits.
func slotKey(slot int) (string, string) {
	k := core.Kind(slot >> 7)
	if k == core.KindReference {
		return k.String(), fmt.Sprintf("0x%X", slot&0x7F)
	}
	return k.String(), fmt.Sprintf("0x%02X", slot&0x7F)
}
