//ff:func feature=live type=tracker control=iteration dimension=1
//ff:what TestPartitions — 한도 없음, 표에 다 들어감, 나눔, 아주 작은 한도
package tally

import "testing"

func TestPartitions(t *testing.T) {
	cases := []struct {
		n             uint64
		memory        int64
		parts, marks_ uint64
	}{
		{1000, 0, 1, 1000},
		{10, 48 * markSize, 1, 10},
		{24, 48 * markSize, 1, 24},
		{25, 48 * markSize, 2, 24},
		{240_000_000, 2 << 30, 6, 44739242},
		{5, 1, 6, 1},
	}
	for _, c := range cases {
		parts, marks := partitions(c.n, c.memory)
		if parts != c.parts || marks != c.marks_ {
			t.Fatalf("%+v: got %d %d", c, parts, marks)
		}
	}
}
