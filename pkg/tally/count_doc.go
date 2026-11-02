//ff:func feature=count type=counter control=sequence
//ff:what countDoc — 살아 있는 문서 하나를 풀어 그 Issuer의 표에 센다
package tally

import (
	"context"
	"fmt"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/pack"
)

// countDoc reads the record of e and counts it into r.
func countDoc(ctx context.Context, src Source, e pack.InventoryEntry, spec codecspec.Spec, set string, r *Report) error {
	b, err := src.Record(ctx, e.Loc)
	if err != nil {
		return fmt.Errorf("%s: %w", e.Path, err)
	}
	unopened, err := CountDocument(b, spec, r.counts(issuerOf(e.Path, set)))
	if err != nil {
		return fmt.Errorf("%s: %w", e.Path, err)
	}
	r.Unopened += unopened
	return nil
}
