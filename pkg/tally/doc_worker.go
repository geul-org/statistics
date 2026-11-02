//ff:func feature=count type=counter control=iteration dimension=1
//ff:what docWorker — 일감의 문서를 하나씩 세어 자기 Report에 담는다, 오류나 취소에 멈춘다
package tally

import (
	"context"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/pack"
)

// docWorker counts jobs into r until they end or ctx does. An error cancels ctx.
func docWorker(ctx context.Context, cancel context.CancelCauseFunc, jobs <-chan pack.InventoryEntry, src Source, spec codecspec.Spec, set string, r *Report) {
	for e := range jobs {
		if ctx.Err() != nil {
			return
		}
		if err := countDoc(ctx, src, e, spec, set, r); err != nil {
			cancel(err)
			return
		}
	}
}
