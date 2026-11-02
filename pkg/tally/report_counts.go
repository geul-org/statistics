//ff:func feature=report type=helper control=sequence
//ff:what Report.counts — Issuer의 표, 없으면 만든다
package tally

// counts returns issuer's table.
func (r *Report) counts(issuer string) *Counts {
	c, ok := r.Issuers[issuer]
	if !ok {
		c = &Counts{}
		r.Issuers[issuer] = c
	}
	return c
}
