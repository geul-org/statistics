//ff:func feature=live type=tracker control=iteration dimension=1
//ff:what findDead — 경로 해시를 칸으로 나눠 칸마다 인벤토리를 다시 읽고 죽은 항목 비트를 만든다
package tally

import "context"

// findDead finds, from the inventories alone, every step a later step
// replaces. pack.Reader.State would hold every path at once; here one table
// of bounded size serves one partition of the paths at a time, and what stays
// is one bit per step.
func findDead(ctx context.Context, src Source, packs []uint32, memory int64) (deadSet, error) {
	steps, tracked, err := countSteps(ctx, src, packs)
	if err != nil {
		return nil, err
	}
	dead := newDeadSet(steps)
	parts, capacity := partitions(tracked, memory)
	marks := make([]mark, capacity)
	for part := range parts {
		n, err := fillMarks(ctx, src, packs, part, parts, marks)
		if err != nil {
			return nil, err
		}
		markDead(marks[:n], dead)
	}
	return dead, nil
}
