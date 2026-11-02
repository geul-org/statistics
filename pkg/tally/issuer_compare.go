//ff:func feature=report type=helper control=sequence
//ff:what issuerCompare — Issuer ID를 폭이 좁은 것부터, 같은 폭은 값 순으로 비교한다
package tally

import "cmp"

// issuerCompare orders canonical Issuer IDs by width, then value.
func issuerCompare(a, b string) int {
	return cmp.Or(cmp.Compare(len(a), len(b)), cmp.Compare(a, b))
}
