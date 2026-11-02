//ff:func feature=report type=writer control=iteration dimension=1
//ff:what writeTSV — 머리 줄과 Issuer·Kind·Codec 순의 행을 TSV로 쓴다, 0인 칸은 뺀다
package tally

import (
	"bufio"
	"fmt"
	"io"
	"maps"
	"slices"
)

// writeTSV writes the columns issuer, kind, codec and count. Issuers sort
// short IDs first, as hex digits are the width.
func writeTSV(w io.Writer, r *Report) error {
	out := bufio.NewWriter(w)
	fmt.Fprintln(out, "issuer\tkind\tcodec\tcount")
	issuers := slices.SortedFunc(maps.Keys(r.Issuers), issuerCompare)
	for _, issuer := range issuers {
		writeIssuer(out, issuer, r.Issuers[issuer])
	}
	return out.Flush()
}
