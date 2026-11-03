//ff:func feature=archive type=resolver control=sequence
//ff:what TestOpenArchives — pack set을 연다, pack이 아닌 경로와 깨진 매니페스트는 오류
package tally

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestOpenArchives(t *testing.T) {
	ctx := context.Background()
	wk := specArchive(t)
	archives, err := openArchives(ctx, []string{wk})
	if err != nil || len(archives) != 1 || len(archives[0].Sets()) != 1 {
		t.Fatalf("%v %v", archives, err)
	}
	if _, err := openArchives(ctx, []string{t.TempDir()}); err == nil || !strings.Contains(err.Error(), "no pack set") {
		t.Fatalf("got %v", err)
	}
	if err := os.WriteFile(filepath.Join(wk, "0x00000000", "latest.txt"), []byte("1 00\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := openArchives(ctx, []string{wk}); err == nil {
		t.Fatal("broken latest.txt: want error")
	}
}
