//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestRun — 사용법 오류, 잘못된 플래그, 집계 실패, 성공과 한도 적용
package main

import (
	"bytes"
	"context"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/pack"
	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestRun(t *testing.T) {
	procs := runtime.GOMAXPROCS(0)
	limit := debug.SetMemoryLimit(-1)
	t.Cleanup(func() { runtime.GOMAXPROCS(procs); debug.SetMemoryLimit(limit) })
	var out, errs bytes.Buffer
	for _, args := range [][]string{{}, {"-memory", "-1", "x"}, {"-threads", "0", "x"}, {"-nope"}} {
		if got := run(args, &out, &errs); got != 2 {
			t.Fatalf("%v: %d", args, got)
		}
	}
	if got := run([]string{t.TempDir()}, &out, &errs); got != 1 || !strings.Contains(errs.String(), "geulstatic: ") {
		t.Fatalf("%d %s", got, errs.String())
	}
	root := t.TempDir()
	target := core.TargetWord(core.KindReference, core.RefWikidataQID<<3)
	id, err := core.EncodeDocID([]core.Word{0}, []core.Word{0, 0, 0, target})
	if err != nil {
		t.Fatal(err)
	}
	var in []core.Word
	for _, r := range []core.Record{core.EncodeSpecVersion(core.SpecVersion), id, {Kind: core.KindMeta, Codec: core.MetaBody},
		core.EncodeDeclare(core.Declare{EdgeID: 1, Target: target})} {
		ws, err := core.EncodeRecord(r)
		if err != nil {
			t.Fatal(err)
		}
		in = append(in, ws...)
	}
	doc := core.PutWords(core.WrapDocument(append(in, 0x180F)))
	if _, err := (resolve.Dir{Root: root}).PublishCodec(doc); err != nil {
		t.Fatal(err)
	}
	wk := filepath.Join(t.TempDir(), "geul")
	if _, err := pack.Split(context.Background(), root, wk, pack.BuildOptions{}); err != nil {
		t.Fatal(err)
	}
	out.Reset()
	if got := run([]string{"-memory", "1000000000", "-threads", "2", wk}, &out, &errs); got != 0 {
		t.Fatalf("%d %s", got, errs.String())
	}
	if !strings.Contains(out.String(), "0x0000\tEDGE\t0x1B\t1\n") {
		t.Fatalf("got %s", out.String())
	}
	if runtime.GOMAXPROCS(0) != 2 || debug.SetMemoryLimit(-1) != 1000000000 {
		t.Fatal("limits not applied")
	}
	if got := run([]string{wk}, &out, &errs); got != 0 {
		t.Fatalf("no limits: %d", got)
	}
}
