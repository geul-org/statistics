//ff:func feature=count type=counter control=sequence
//ff:what TestDocWorker — 일감을 다 세고 끝난다, 오류는 취소로 알린다, 취소되면 멈춘다
package tally

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/pack"
)

func TestDocWorker(t *testing.T) {
	end := slotOf(core.KindMeta, core.MetaEnd)
	s := fakeSource{docs: map[uint64][]byte{0: docOf(t)}, err: errors.New("broken")}
	jobs := make(chan pack.InventoryEntry, 2)
	jobs <- pack.InventoryEntry{Path: "a.geul"}
	jobs <- pack.InventoryEntry{Path: "b.geul"}
	close(jobs)
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)
	r := newReport()
	docWorker(ctx, cancel, jobs, s, testSpec(), "0x0001", r)
	if ctx.Err() != nil || r.Issuers["0x0001"][end] != 2 {
		t.Fatalf("%v %+v", ctx.Err(), r)
	}
	s.recFail = true
	jobs = make(chan pack.InventoryEntry, 1)
	jobs <- pack.InventoryEntry{Path: "a.geul"}
	docWorker(ctx, cancel, jobs, s, testSpec(), "0x0001", newReport())
	if !errors.Is(context.Cause(ctx), s.err) {
		t.Fatalf("got %v", context.Cause(ctx))
	}
	jobs = make(chan pack.InventoryEntry, 1)
	jobs <- pack.InventoryEntry{Path: "a.geul"}
	r = newReport()
	docWorker(ctx, cancel, jobs, s, testSpec(), "0x0001", r)
	if len(jobs) != 0 || len(r.Issuers) != 0 {
		t.Fatalf("cancelled: %d %+v", len(jobs), r)
	}
}
