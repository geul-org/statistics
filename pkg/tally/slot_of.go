//ff:func feature=count type=helper control=sequence
//ff:what slotOf — 레코드의 칸 번호: TAG의 9비트 코드(w>>4, TAG ID 0x000L~0x1FFL), LLLL만 버리고 묶지 않는다
package tally

import "github.com/geul-org/gogeul/pkg/core"

// slotOf is the slot of a record: the 9-bit code of its TAG WORD `000K KKCC
// CCCC LLLL` (w>>4, TAG ID 0x000L~0x1FFL), the start of its Kind's zone plus
// its Record Codec. Only LLLL is dropped: a Reference keeps its Quantifier,
// and the slot markers and pointers are codes of their own.
func slotOf(k core.Kind, codec uint16) int {
	return int(core.TargetWord(k, codec) >> 4)
}
