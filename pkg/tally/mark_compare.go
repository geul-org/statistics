//ff:func feature=live type=tracker control=sequence
//ff:what markCompare — 같은 경로끼리 모이고 그 안에서 재생 순서가 되게 비교한다
package tally

import "cmp"

// markCompare sorts by path hash, then replay order.
func markCompare(a, b mark) int {
	return cmp.Or(cmp.Compare(a.h0, b.h0), cmp.Compare(a.h1, b.h1), cmp.Compare(a.seq, b.seq))
}
