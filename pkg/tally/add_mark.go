//ff:func feature=live type=tracker control=sequence
//ff:what addMark — 문서 경로이고 칸 part에 들면 표에 한 줄 더한다
package tally

// addMark appends the mark of path at step seq to marks[:*used] when path is a
// document path of partition part.
func addMark(path string, seq, part, parts uint64, marks []mark, used *int) error {
	if !isDocPath(path) {
		return nil
	}
	m := markOf(path, seq)
	if m.h0%parts != part {
		return nil
	}
	if *used == len(marks) {
		return errMarksFull
	}
	marks[*used] = m
	*used++
	return nil
}
