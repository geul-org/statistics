//ff:func feature=live type=tracker control=sequence
//ff:what deadSet.mark — 순번의 비트를 세운다
package tally

// mark sets seq's bit.
func (d deadSet) mark(seq uint64) {
	step := uint32(seq)
	d[seq>>32][step/64] |= 1 << (step % 64)
}
