//ff:func feature=archive type=helper control=sequence
//ff:what writeGen — 시험용: pack set에 세대 하나를 쓴다
package tally

import (
	"context"
	"maps"
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

// writeGen adds a generation putting puts and deleting dels.
func writeGen(t *testing.T, dir string, puts map[string][]byte, dels ...string) {
	t.Helper()
	b := pack.Batch{Put: slices.Sorted(maps.Keys(puts)), Delete: dels, Get: func(p string) ([]byte, error) { return puts[p], nil }}
	if _, err := pack.Write(context.Background(), dir, b, pack.BuildOptions{}); err != nil {
		t.Fatal(err)
	}
}
