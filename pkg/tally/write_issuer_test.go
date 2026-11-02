//ff:func feature=report type=writer control=sequence
//ff:what TestWriteIssuer — 0이 아닌 칸만 쓴다
package tally

import (
	"bytes"
	"testing"
)

func TestWriteIssuer(t *testing.T) {
	var out bytes.Buffer
	writeIssuer(&out, "0x0001", &Counts{2<<7 | 0x0A: 9})
	if out.String() != "0x0001\tEDGE\t0x0A\t9\n" {
		t.Fatalf("got %q", out.String())
	}
}
