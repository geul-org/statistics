//ff:func feature=count type=helper control=iteration dimension=1
//ff:what nonZero — 시험 메시지용: 표의 0이 아닌 칸
package tally

// nonZero lists c's used slots.
func nonZero(c Counts) map[int]uint64 {
	out := map[int]uint64{}
	for i, n := range c {
		if n != 0 {
			out[i] = n
		}
	}
	return out
}
