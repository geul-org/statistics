//ff:func feature=archive type=resolver control=iteration dimension=1
//ff:what specTargets — 규격 문서가 있을 수 있는 대상 WORD 전부: Reference 15개, 나머지 Kind의 직접 코드
package tally

import "github.com/geul-org/gogeul/pkg/core"

// specTargets lists every direct codec's target word. Codec 0000 is reserved
// and has no spec, and ISSUER codes are not looked up (codecspec.Spec.Of).
func specTargets() []core.Word {
	var out []core.Word
	for c := uint8(1); c < 16; c++ {
		out = append(out, core.TargetWord(core.KindReference, c<<3))
	}
	for c := uint8(0); c < 0x7C; c++ {
		out = append(out, core.TargetWord(core.KindValue, c), core.TargetWord(core.KindEdge, c), core.TargetWord(core.KindMeta, c))
	}
	return out
}
