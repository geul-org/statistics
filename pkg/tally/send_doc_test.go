//ff:func feature=count type=helper control=iteration dimension=1
//ff:what TestSendDoc — h/ 전용은 건수만, 죽은 항목·삭제·문서 아닌 경로는 버리고, 살아 있는 문서만 보낸다, 작업자가 멈추면 오류
package tally

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestSendDoc(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)
	jobs := make(chan pack.InventoryEntry, 1)
	r := newReport()
	live := pack.InventoryEntry{Path: "0x00000001/docs/0x000000000000000A/0x00000001.geul"}
	skip := []struct {
		e    pack.InventoryEntry
		dead bool
	}{
		{live, true},
		{pack.InventoryEntry{Path: "0x00000001/docs/0x000000000000000A/0x00000001.geul", Deleted: true}, false},
		{pack.InventoryEntry{Path: "0x00000001/hashes/0xAB.geul", Deleted: true}, false},
		{pack.InventoryEntry{Path: "keys/a.pub"}, false},
		{pack.InventoryEntry{Path: "0x00000001/docs/versions.geul"}, false},
		{pack.InventoryEntry{Hashed: true}, false},
	}
	for _, c := range skip {
		if err := sendDoc(ctx, c.e, c.dead, jobs, r); err != nil || len(jobs) != 0 {
			t.Fatalf("%+v: %v %d", c, err, len(jobs))
		}
	}
	if r.HashOnly != 1 {
		t.Fatalf("HashOnly %d", r.HashOnly)
	}
	if err := sendDoc(ctx, live, false, jobs, r); err != nil || len(jobs) != 1 {
		t.Fatalf("%v %d", err, len(jobs))
	}
	boom := errors.New("worker stopped")
	cancel(boom)
	if err := sendDoc(ctx, live, false, jobs, r); !errors.Is(err, boom) {
		t.Fatalf("got %v", err)
	}
}
