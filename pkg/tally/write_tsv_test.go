//ff:func feature=report type=writer control=sequence
//ff:what TestWriteTSV — 머리 줄, Issuer 폭 순, 칸 순, 0인 칸 없음, 쓰기 오류
package tally

import (
	"bytes"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestWriteTSV(t *testing.T) {
	r := newReport()
	r.counts("0x00010000")[slotOf(core.KindEdge, core.EdgeTriple)] = 5
	c := r.counts("0x0000")
	c[slotOf(core.KindMeta, core.MetaEnd)] = 2
	c[slotOf(core.KindReference, 0x20)] = 3
	var out bytes.Buffer
	if err := writeTSV(&out, r); err != nil {
		t.Fatal(err)
	}
	want := "issuer\tkind\tcodec\tcount\n" +
		"0x0000\tREFERENCE\t0x4\t3\n" +
		"0x0000\tMETA\t0x0F\t2\n" +
		"0x00010000\tEDGE\t0x0A\t5\n"
	if out.String() != want {
		t.Fatalf("got %q", out.String())
	}
	if err := writeTSV(failWriter{}, r); err == nil {
		t.Fatal("want write error")
	}
}
