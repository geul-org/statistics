//ff:func feature=archive type=resolver control=iteration dimension=1
//ff:what openArchives — 경로마다 pack 아카이브를 연다, pack set이 없는 경로는 오류다
package tally

import (
	"context"
	"fmt"

	"github.com/geul-org/gogeul/pkg/pack"
)

// openArchives opens each path: a `.well-known/geul/` directory, or one set's
// directory. Opening reads manifests and layer roots only.
func openArchives(ctx context.Context, paths []string) ([]*pack.Archive, error) {
	var out []*pack.Archive
	for _, p := range paths {
		if !pack.IsPacked(p) {
			return nil, fmt.Errorf("%s holds no pack set", p)
		}
		a, err := pack.OpenDir(ctx, p, pack.BuildOptions{})
		if err != nil {
			return nil, fmt.Errorf("%s: %w", p, err)
		}
		out = append(out, a)
	}
	return out, nil
}
