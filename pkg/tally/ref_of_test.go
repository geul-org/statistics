//ff:func feature=count type=helper control=sequence
//ff:what refOf — 시험용: 묶음 머리 코덱(Identified)과 Quantifier q의 Reference, 머리 0이면 자리 표지 C q
package tally

import "github.com/geul-org/gogeul/pkg/core"

// refOf builds a Reference of the 5Q scheme whose Identified Codec is head
// (core.CodecLocal, core.CodecWikidataQID, …) and Quantifier q: C = head + q
// (§3.15). Head 0 gives the slot markers C 0, 1, 2.
func refOf(head, q uint16, id ...core.Word) core.Record {
	return core.Record{Kind: core.KindReference, Codec: head + q, Payload: id}
}
