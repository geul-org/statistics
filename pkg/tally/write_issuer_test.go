//ff:func feature=report type=writer control=sequence
//ff:what TestWriteIssuer — 0이 아닌 칸만 쓴다
package tally

import (
	"bytes"
	"testing"
)

func TestWriteIssuer(t *testing.T) {
	var out bytes.Buffer
	writeIssuer(&out, "0x00000001", &Counts{0x0AA: 9})
	if out.String() != "0x00000001\tEDGE\t0x0AAL\t9\n" {
		t.Fatalf("got %q", out.String())
	}
}
