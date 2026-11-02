//ff:func feature=archive type=counter control=iteration dimension=1
//ff:what CountSet — pack set 하나: 죽은 항목을 찾고 pack 차례로 살아 있는 문서를 센다
package tally

import (
	"context"
	"runtime"

	"github.com/geul-org/gogeul/pkg/codecspec"
)

// CountSet counts the live documents of src, whose set belongs to Issuer set.
// Finding the dead steps takes one thread; counting takes o.Threads, set here
// for the whole process, GC and zstd included.
func CountSet(ctx context.Context, src Source, spec codecspec.Spec, set string, o Options) (*Report, error) {
	packs := src.DataPacks()
	dead, err := findDead(ctx, src, packs, o.Memory)
	if err != nil {
		return nil, err
	}
	runtime.GOMAXPROCS(o.Threads)
	total := newReport()
	for i, obj := range packs {
		if err := countPack(ctx, src, i, obj, dead, spec, set, o.Threads, total); err != nil {
			return nil, err
		}
	}
	return total, nil
}
