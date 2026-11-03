//ff:func feature=archive type=counter control=sequence
//ff:what TestCountArchives — 아카이브들의 set을 합친다, 세다 깨지면 오류
package tally

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCountArchives(t *testing.T) {
	ctx := context.Background()
	dir := filepath.Join(t.TempDir(), "0x00000002")
	writeGen(t, dir, map[string][]byte{"0x00000002/docs/0x0000000000000001/0x00000001.geul": docOf(t)})
	archives, err := openArchives(ctx, []string{dir, dir})
	if err != nil {
		t.Fatal(err)
	}
	r, err := countArchives(ctx, archives, testSpec(), Options{Threads: 1})
	if err != nil {
		t.Fatal(err)
	}
	if r.Issuers["0x00000002"][slotOf(core.KindMeta, core.MetaEnd)] != 2 {
		t.Fatalf("got %+v", r)
	}
	if err := os.Remove(filepath.Join(dir, "latest.txt")); err != nil {
		t.Fatal(err)
	}
	if _, err := countArchives(ctx, archives, testSpec(), Options{Threads: 1}); err == nil {
		t.Fatal("want error")
	}
}
