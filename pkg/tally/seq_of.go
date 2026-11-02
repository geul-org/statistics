//ff:func feature=live type=helper control=sequence
//ff:what seqOf — pack 위치와 인벤토리 단계로 재생 순번을 만든다
package tally

// seqOf orders steps as replay does: pack position, then step in the pack.
func seqOf(i int, step uint32) uint64 {
	return uint64(i)<<32 | uint64(step)
}
