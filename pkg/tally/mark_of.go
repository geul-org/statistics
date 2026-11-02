//ff:func feature=live type=tracker control=sequence
//ff:what markOf — 경로와 순번으로 표의 한 줄을 만든다
package tally

import (
	"crypto/sha256"
	"encoding/binary"
)

// markOf hashes path for the mark of step seq.
func markOf(path string, seq uint64) mark {
	s := sha256.Sum256([]byte(path))
	return mark{h0: binary.BigEndian.Uint64(s[:8]), h1: binary.BigEndian.Uint64(s[8:16]), seq: seq}
}
