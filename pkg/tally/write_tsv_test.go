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
	r.counts("0x0000000100000000")[slotOf(core.KindEdge, core.EdgeTriple)] = 5
	c := r.counts("0x00000000")
	c[slotOf(core.KindMeta, core.MetaEnd)] = 2
	c[slotOf(core.KindReference, core.CodecWikidataQID)] = 3
	var out bytes.Buffer
	if err := writeTSV(&out, r); err != nil {
		t.Fatal(err)
	}
	want := "issuer\tkind\tcodec\tcount\n" +
		"0x00000000\tREFERENCE\t0x014L\t3\n" +
		"0x00000000\tMETA\t0x1C4L\t2\n" +
		"0x0000000100000000\tEDGE\t0x0AAL\t5\n"
	if out.String() != want {
		t.Fatalf("got %q", out.String())
	}
	if err := writeTSV(failWriter{}, r); err == nil {
		t.Fatal("want write error")
	}
}
