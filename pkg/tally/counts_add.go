//ff:func feature=report type=helper control=iteration dimension=1
//ff:what Counts.add — 표를 칸별로 더한다
package tally

// add adds o to c.
func (c *Counts) add(o *Counts) {
	for i, n := range o {
		c[i] += n
	}
}
