//ff:func feature=count type=helper control=sequence
//ff:what slotOf — Kind와 TAG의 Codec 7비트로 칸 번호를 만든다, Reference는 하위 3비트를 버린다
package tally

import "github.com/geul-org/gogeul/pkg/core"

// slotOf is the slot of a record: KK then CCCCCCC, or KK then CCCC for a
// Reference, whose low 3 bits are a Quantifier or a slot marker number.
func slotOf(k core.Kind, codec uint8) int {
	if k == core.KindReference {
		codec >>= 3
	}
	return int(k)<<7 | int(codec)
}
