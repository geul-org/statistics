//ff:func feature=report type=helper control=iteration dimension=1
//ff:what Report.add — 다른 Report를 더한다
package tally

// add adds o to r.
func (r *Report) add(o *Report) {
	r.Unopened += o.Unopened
	r.HashOnly += o.HashOnly
	for issuer, c := range o.Issuers {
		r.counts(issuer).add(c)
	}
}
