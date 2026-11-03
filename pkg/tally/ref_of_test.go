//ff:func feature=count type=helper control=sequence
//ff:what refOf — 시험용: 체계 scheme(C/5)와 Quantifier q의 Reference, 체계 0이면 자리 표지 C q
package tally

import "github.com/geul-org/gogeul/pkg/core"

// refOf builds a Reference of scheme (C/5) and Quantifier q: C = scheme*5 + q
// (§3.1). Scheme 0 gives the slot markers C 0, 1, 2.
func refOf(scheme, q uint16, id ...core.Word) core.Record {
	return core.Record{Kind: core.KindReference, Codec: scheme*5 + q, Payload: id}
}
