//ff:func feature=count type=counter control=sequence
//ff:what TestCountDoc — 문서를 Issuer 표에 센다, 읽기 오류와 깨진 문서는 경로를 달아 돌려준다
package tally

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/pack"
)

func TestCountDoc(t *testing.T) {
	ctx := context.Background()
	e := pack.InventoryEntry{Path: "0x00000001/docs/0x0000000000000001/0x00000001.geul", Loc: pack.Loc{Off: 1}}
	s := fakeSource{docs: map[uint64][]byte{1: docOf(t, core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("ab"))), 2: {1}}}
	r := newReport()
	if err := countDoc(ctx, s, e, testSpec(), "0x00000000", r); err != nil {
		t.Fatal(err)
	}
	if r.Unopened != 1 || r.Issuers["0x00000001"][slotOf(core.KindMeta, core.MetaEnd)] != 1 {
		t.Fatalf("got %+v", r)
	}
	bad := e
	bad.Loc.Off = 2
	if err := countDoc(ctx, s, bad, testSpec(), "0x00000000", r); err == nil || !strings.Contains(err.Error(), e.Path) {
		t.Fatalf("got %v", err)
	}
	s.recFail, s.err = true, errors.New("broken")
	if err := countDoc(ctx, s, e, testSpec(), "0x00000000", r); !errors.Is(err, s.err) || !strings.Contains(err.Error(), e.Path) {
		t.Fatalf("got %v", err)
	}
}
