//ff:type feature=live type=model
//ff:what mark — 경로 추적 표의 한 줄: 경로 해시 128비트와 항목 순번
package tally

// markSize is the bytes of one mark in the table.
const markSize = 24

// mark is one inventory step that names a document path: the first 128 bits
// of the path's SHA-256, and the step's place in replay order (seqOf).
type mark struct {
	h0, h1 uint64
	seq    uint64
}
