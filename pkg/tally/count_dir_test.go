//ff:func feature=archive type=counter control=sequence
//ff:what TestCountDir — set 폴더의 문서를 폴더 이름의 Issuer로 센다, 없는 폴더는 오류
package tally

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCountDir(t *testing.T) {
	ctx := context.Background()
	dir := filepath.Join(t.TempDir(), "0x00000000")
	writeGen(t, dir, map[string][]byte{"issuers/0x00000001.geul": docOf(t)})
	r, err := countDir(ctx, dir, testSpec(), Options{Threads: 1})
	if err != nil {
		t.Fatal(err)
	}
	if r.Issuers["0x00000000"][slotOf(core.KindMeta, core.MetaEnd)] != 1 {
		t.Fatalf("got %+v", r)
	}
	if _, err := countDir(ctx, filepath.Join(dir, "none"), testSpec(), Options{Threads: 1}); err == nil {
		t.Fatal("want error")
	}
}
