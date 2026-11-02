//ff:func feature=live type=tracker control=sequence
//ff:what deadSet.has — pack 위치와 단계의 비트를 읽는다
package tally

// has reports whether step of pack position i is dead.
func (d deadSet) has(i int, step uint32) bool {
	return d[i][step/64]&(1<<(step%64)) != 0
}
