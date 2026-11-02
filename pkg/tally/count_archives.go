//ff:func feature=archive type=counter control=iteration dimension=2
//ff:what countArchives — 아카이브의 set을 차례로 세어 합친다
package tally

import (
	"context"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/pack"
)

// countArchives counts one set at a time, so the limits hold per set.
func countArchives(ctx context.Context, archives []*pack.Archive, spec codecspec.Spec, o Options) (*Report, error) {
	total := newReport()
	for _, a := range archives {
		for _, s := range a.Sets() {
			r, err := countDir(ctx, s.Dir, spec, o)
			if err != nil {
				return nil, err
			}
			total.add(r)
		}
	}
	return total, nil
}
