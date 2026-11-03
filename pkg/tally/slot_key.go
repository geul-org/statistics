//ff:func feature=report type=helper control=sequence
//ff:what slotKey — 칸 번호를 TSV의 kind와 codec 글자(TAG ID 0x000L 표기)로 되돌린다
package tally

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// slotKey names a slot: the Kind's spec name, read from KKK, and the code as
// a TAG ID `0x000L` (three hex digits and L for the Length bits left out).
func slotKey(slot int) (string, string) {
	return core.KindOf(core.Word(slot) << 4).String(), fmt.Sprintf("0x%03XL", slot)
}
