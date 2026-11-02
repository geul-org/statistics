//ff:func feature=count type=counter control=iteration dimension=1
//ff:what TestCountPack — 살아 있는 문서만 합쳐 세고, 문서 오류와 인벤토리 오류를 돌려준다
package tally

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCountPack(t *testing.T) {
	ctx := context.Background()
	s := twoPackSource()
	s.docs = map[uint64][]byte{}
	for off := range uint64(10) {
		s.docs[off] = docOf(t)
	}
	dead, err := findDead(ctx, s, s.packs, 0)
	if err != nil {
		t.Fatal(err)
	}
	total := newReport()
	if err := countPack(ctx, s, 1, s.packs[1], dead, testSpec(), "0x0001", 3, total); err != nil {
		t.Fatal(err)
	}
	if got := total.Issuers["0x0001"][slotOf(core.KindMeta, core.MetaEnd)]; got != 2 || total.HashOnly != 1 {
		t.Fatalf("got %d %+v", got, total) // b and d
	}
	s.recFail, s.err = true, errors.New("broken record")
	if err := countPack(ctx, s, 0, s.packs[0], dead, testSpec(), "0x0001", 2, newReport()); !errors.Is(err, s.err) {
		t.Fatalf("got %v", err)
	}
	s.recFail, s.invCall, s.invFail = false, new(int), 1
	if err := countPack(ctx, s, 0, s.packs[0], dead, testSpec(), "0x0001", 2, newReport()); !errors.Is(err, s.err) {
		t.Fatalf("got %v", err)
	}
}
