//ff:func feature=live type=tracker control=iteration dimension=1
//ff:what newDeadSet — pack별 단계 수에 맞는 빈 비트 표
package tally

// newDeadSet sizes the bits for steps, the step count of each pack.
func newDeadSet(steps []uint64) deadSet {
	d := make(deadSet, len(steps))
	for i, n := range steps {
		d[i] = make([]uint64, (n+63)/64)
	}
	return d
}
