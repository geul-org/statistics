//ff:func feature=archive type=counter control=sequence
//ff:what countDir — set 폴더 하나를 열어 CountSet을 돌린다
package tally

import (
	"context"
	"fmt"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/pack"
)

// countDir counts the set in dir with a Reader of its own, its latest generation.
func countDir(ctx context.Context, dir string, spec codecspec.Spec, o Options) (*Report, error) {
	r, err := pack.Open(ctx, pack.FileStore{Root: dir}, pack.Options{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", dir, err)
	}
	return CountSet(ctx, r, spec, setIssuer(dir), o)
}
