//ff:func feature=count type=counter control=iteration dimension=1
//ff:what countPack — pack 하나의 인벤토리를 흘려 살아 있는 문서를 작업자들에게 나눠 세고 합친다
package tally

import (
	"context"
	"sync"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/pack"
)

// jobsPerThread is how many documents wait per worker. Few are enough, as a
// job holds only a path and a location.
const jobsPerThread = 16

// countPack counts pack position i, object obj, into total. One goroutine
// walks the inventory, which stays in memory until it ends, so only one is
// open at a time; threads workers read and count the documents one by one.
func countPack(ctx context.Context, src Source, i int, obj uint32, dead deadSet, spec codecspec.Spec, set string, threads int, total *Report) error {
	ctx, cancel := context.WithCancelCause(ctx)
	defer cancel(nil)
	jobs := make(chan pack.InventoryEntry, threads*jobsPerThread)
	reports := make([]*Report, threads)
	var wg sync.WaitGroup
	for t := range reports {
		reports[t] = newReport()
		wg.Go(func() { docWorker(ctx, cancel, jobs, src, spec, set, reports[t]) })
	}
	step := uint32(0)
	err := src.PackInventory(ctx, obj, func(e pack.InventoryEntry) error {
		step++
		return sendDoc(ctx, e, dead.has(i, step-1), jobs, total)
	})
	close(jobs)
	wg.Wait()
	for _, r := range reports {
		total.add(r)
	}
	return firstError(context.Cause(ctx), err)
}
