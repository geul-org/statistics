//ff:func feature=archive type=command control=iteration dimension=1
//ff:what TestRun — 끝에서 끝까지: TSV 행, 세지 못한 건수 알림, 열기·규격·세기·쓰기 오류
package tally

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRun(t *testing.T) {
	ctx := context.Background()
	wk := specArchive(t)
	set := filepath.Join(wk, "0x0001")
	qid := refOf(core.RefWikidataQID, 1, 42)
	h := docOf(t, qid)
	ws, _ := core.ReadWords(h)
	sum, err := core.ContentHash(ws)
	if err != nil {
		t.Fatal(err)
	}
	writeGen(t, set, map[string][]byte{
		"0x0001/0x0000000000000001/0x00000001.geul":   docOf(t, qid, refOf(core.RefLocal, 0, 7)),
		"h/0x" + hex.EncodeToString(sum[:]) + ".geul": h,
	})
	var out, notice bytes.Buffer
	if err := Run(ctx, []string{wk}, Options{Threads: 2}, &out, &notice); err != nil {
		t.Fatal(err)
	}
	for _, row := range []string{"0x0001\tREFERENCE\t0x4\t1\n", "0x0001\tREFERENCE\t0x1\t1\n", "0x0001\tMETA\t0x0F\t1\n"} {
		if !strings.Contains(out.String(), row) {
			t.Fatalf("no %q in\n%s", row, out.String())
		}
	}
	if !strings.HasPrefix(out.String(), "issuer\tkind\tcodec\tcount\n0x0000\t") {
		t.Fatalf("got %s", out.String())
	}
	if n := notice.String(); !strings.Contains(n, "payloads have no usable") || !strings.Contains(n, "1 h/-only records") {
		t.Fatalf("notice %q", n)
	}
	if err := Run(ctx, []string{t.TempDir()}, Options{Threads: 1}, &out, &notice); err == nil {
		t.Fatal("open: want error")
	}
	if err := Run(ctx, []string{set}, Options{Threads: 1}, &out, &notice); !errors.Is(err, errNoSpec) {
		t.Fatalf("spec: %v", err)
	}
	if err := Run(ctx, []string{wk}, Options{Threads: 1}, failWriter{}, &notice); err == nil {
		t.Fatal("write: want error")
	}
	objects, _ := filepath.Glob(filepath.Join(set, "o", "*.gpk"))
	for _, o := range objects {
		if err := os.Truncate(o, 8); err != nil {
			t.Fatal(err)
		}
	}
	if err := Run(ctx, []string{wk}, Options{Threads: 1}, &out, &notice); err == nil {
		t.Fatal("count: want error")
	}
}
