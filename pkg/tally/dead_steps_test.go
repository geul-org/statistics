//ff:func feature=live type=helper control=iteration dimension=1
//ff:what checkDeadSteps — twoPackSource의 죽은 단계가 손으로 정한 대로인지 본다
package tally

import (
	"testing"
)

// checkDeadSteps wants exactly the replaced steps dead: in pack 0 a, b, c and
// d's first records, in pack 1 the deletion of d, which d's new record replaces.
func checkDeadSteps(t *testing.T, d deadSet) {
	t.Helper()
	dead := []uint64{seqOf(0, 0), seqOf(0, 1), seqOf(0, 2), seqOf(0, 3), seqOf(1, 2)}
	for _, s := range []uint64{seqOf(0, 4), seqOf(1, 0), seqOf(1, 1), seqOf(1, 3), seqOf(1, 4), seqOf(1, 5), seqOf(1, 6)} {
		if d.has(int(s>>32), uint32(s)) {
			t.Fatalf("seq %x: dead", s)
		}
	}
	for _, s := range dead {
		if !d.has(int(s>>32), uint32(s)) {
			t.Fatalf("seq %x: alive", s)
		}
	}
}
