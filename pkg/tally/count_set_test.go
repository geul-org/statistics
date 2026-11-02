//ff:func feature=archive type=counter control=iteration dimension=1
//ff:what TestCountSet — 실제 pack 세대로: 덮어쓰기·삭제가 중복되지 않고, h/ 전용은 빠지고, 찾기와 세기의 오류가 돌아온다
package tally

import (
	"context"
	"encoding/hex"
	"errors"
	"runtime"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/pack"
)

func TestCountSet(t *testing.T) {
	procs := runtime.GOMAXPROCS(0)
	t.Cleanup(func() { runtime.GOMAXPROCS(procs) })
	ctx := context.Background()
	dir := t.TempDir()
	pred := refOf(core.RefWikidataPID, 0, 31)
	one := func(obj core.Word) core.Record {
		return tripleOf(t, 1, pred, refOf(core.RefWikidataQID, 0, 42), refOf(core.RefWikidataQID, 0, obj))
	}
	oldA := docOf(t, one(1))
	newA := docOf(t, one(2), one(3))
	gone := docOf(t, one(4), one(5), one(6))
	hashed := docOf(t, one(7))
	ws, _ := core.ReadWords(hashed)
	h, err := core.ContentHash(ws)
	if err != nil {
		t.Fatal(err)
	}
	a := "0x0001/0x0000000000000001/0x00000001.geul"
	b := "0x0001/0x0000000000000002/0x00000001.geul"
	c := "0x0001/0x0000000000000003/0x00000001.geul"
	writeGen(t, dir, map[string][]byte{a: oldA, b: oldA, c: gone, "issuers.geul": docOf(t)})
	writeGen(t, dir, map[string][]byte{a: newA, "h/0x" + hex.EncodeToString(h[:]) + ".geul": hashed}, c)
	r, err := pack.Open(ctx, pack.FileStore{Root: dir}, pack.Options{})
	if err != nil {
		t.Fatal(err)
	}
	for _, memory := range []int64{0, 16 * markSize} {
		got, err := CountSet(ctx, r, testSpec(), "0x0000", Options{Memory: memory, Threads: 2})
		if err != nil {
			t.Fatal(err)
		}
		want := Counts{}
		want[slotOf(core.KindMeta, core.MetaDoc1Word)] = 2 // a@2, b
		want[slotOf(core.KindMeta, core.MetaEnd)] = 2
		want[slotOf(core.KindEdge, core.EdgeTriple)] = 3
		want[slotOf(core.KindReference, core.RefLocal<<3)] = 3
		want[slotOf(core.KindReference, core.RefWikidataPID<<3)] = 3
		want[slotOf(core.KindReference, core.RefWikidataQID<<3)] = 6
		if c := *got.Issuers["0x0001"]; c != want {
			t.Fatalf("got %v want %v", nonZero(c), nonZero(want))
		}
		if runtime.GOMAXPROCS(0) != 2 {
			t.Fatal("threads not applied")
		}
		other := got.Issuers["0x0000"] // issuers.geul belongs to the set's Issuer
		if other[slotOf(core.KindMeta, core.MetaEnd)] != 1 || got.HashOnly != 1 || got.Unopened != 0 {
			t.Fatalf("got %+v", got)
		}
	}
	s := fakeSource{packs: []uint32{1}, invCall: new(int), invFail: 1, err: errors.New("broken")}
	if _, err := CountSet(ctx, s, testSpec(), "0x0000", Options{Threads: 1}); !errors.Is(err, s.err) {
		t.Fatalf("find: %v", err)
	}
	s.invCall, s.invFail = new(int), 3
	if _, err := CountSet(ctx, s, testSpec(), "0x0000", Options{Threads: 1}); !errors.Is(err, s.err) {
		t.Fatalf("count: %v", err)
	}
}
