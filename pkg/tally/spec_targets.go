//ff:func feature=archive type=resolver control=iteration dimension=1
//ff:what specTargets — 규격 문서가 있을 수 있는 대상 WORD 전부: 512 코드 중 자리 표지, 묶음 머리가 아닌 양화, ISSUER 코드를 뺀 것
package tally

import "github.com/geul-org/gogeul/pkg/core"

// specTargets lists the target word (L=0) of every direct code that may have
// a codec document: all 512 codes but the slot markers (C < 3, no spec), the
// Reference Quantifiers other than Identified (a 5Q scheme's spec is its
// Identified head, C - C%5) and the ISSUER codes, which are not looked up
// (codecspec.Spec.Of).
func specTargets() []core.Word {
	var out []core.Word
	for code := core.Word(0); code < 0x200; code++ {
		w := code << 4
		k := core.KindOf(w)
		c := uint16(code) - uint16(core.TargetWord(k, 0)>>4)
		skip := k == core.KindReference && (core.IsSlotMarkerCodec(c) || core.RefIdentified(c) != c)
		if skip || core.IssuerWidth(k, c) != 0 {
			continue
		}
		out = append(out, w)
	}
	return out
}
