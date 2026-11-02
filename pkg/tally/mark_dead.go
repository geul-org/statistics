//ff:func feature=live type=tracker control=iteration dimension=1
//ff:what markDead — 표를 정렬해 같은 경로의 마지막이 아닌 줄을 죽은 항목으로 표시한다
package tally

import "slices"

// markDead sorts marks and marks every step that a later step of the same
// path replaces. A deletion is such a later step, and is itself never counted.
func markDead(marks []mark, dead deadSet) {
	slices.SortFunc(marks, markCompare)
	for i := 0; i+1 < len(marks); i++ {
		if marks[i].h0 == marks[i+1].h0 && marks[i].h1 == marks[i+1].h1 {
			dead.mark(marks[i].seq)
		}
	}
}
