//ff:func feature=report type=writer control=iteration dimension=1
//ff:what writeIssuer — 한 Issuer의 0이 아닌 칸을 칸 번호 순으로 쓴다
package tally

import (
	"fmt"
	"io"
)

// writeIssuer writes c's rows. A write error stays in w for Flush to return.
func writeIssuer(w io.Writer, issuer string, c *Counts) {
	for slot, n := range c {
		if n == 0 {
			continue
		}
		kind, codec := slotKey(slot)
		fmt.Fprintf(w, "%s\t%s\t%s\t%d\n", issuer, kind, codec, n)
	}
}
