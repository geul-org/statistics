//ff:func feature=count type=helper control=sequence
//ff:what refOf — 시험용 Reference 레코드: Codec 4비트, Quantifier 3비트, 대상 ID
package tally

import "github.com/geul-org/gogeul/pkg/core"

// refOf builds a Reference of 4-bit codec cccc and Quantifier q.
func refOf(cccc, q uint8, id ...core.Word) core.Record {
	return core.Record{Kind: core.KindReference, Codec: cccc<<3 | q, Payload: id}
}
