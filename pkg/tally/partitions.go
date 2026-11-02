//ff:func feature=live type=tracker control=sequence
//ff:what partitions — 경로 n개를 메모리 절반에 담도록 나눌 칸 수와 칸당 표 크기를 정한다
package tally

// partitions returns how many passes tracking n steps takes and the marks one
// pass may hold, with half of memory for the table. Paths fall in partitions
// by hash, so a partition is about n/parts: the extra sixteenth absorbs the
// difference. A memory of 0 means one pass of n marks.
func partitions(n uint64, memory int64) (uint64, uint64) {
	capacity := max(uint64(memory)/2/markSize, 1)
	if memory == 0 || n <= capacity {
		return 1, n
	}
	return (n+n/16)/capacity + 1, capacity
}
