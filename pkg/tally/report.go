//ff:type feature=report type=model
//ff:what Report — Issuer별 표와, 세지 못한 것의 건수
package tally

// Report is what counting found. Unopened is how many payloads had no usable
// codec spec: their own TAG is counted, what they embed is not. HashOnly is
// how many h/-only inventory entries were left out.
type Report struct {
	Issuers  map[string]*Counts
	Unopened uint64
	HashOnly uint64
}
