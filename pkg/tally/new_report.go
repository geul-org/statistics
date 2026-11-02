//ff:func feature=report type=model control=sequence
//ff:what newReport — 빈 Report
package tally

// newReport returns an empty Report.
func newReport() *Report {
	return &Report{Issuers: map[string]*Counts{}}
}
