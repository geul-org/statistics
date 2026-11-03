//ff:func feature=archive type=command control=sequence
//ff:what Run — 아카이브를 열고 규격을 찾아 세고, TSV와 세지 못한 건수를 쓴다
package tally

import (
	"context"
	"fmt"
	"io"
)

// Run counts the archives at paths and writes the TSV to out. The same
// document in two archives is counted twice; notes go to notice.
func Run(ctx context.Context, paths []string, o Options, out, notice io.Writer) error {
	archives, err := openArchives(ctx, paths)
	if err != nil {
		return err
	}
	spec, err := specOf(archives)
	if err != nil {
		return err
	}
	r, err := countArchives(ctx, archives, spec, o)
	if err != nil {
		return err
	}
	if r.Unopened > 0 {
		fmt.Fprintf(notice, "statistics: %d payloads have no usable codec spec: their TAGs are counted, not what they embed\n", r.Unopened)
	}
	if r.HashOnly > 0 {
		fmt.Fprintf(notice, "statistics: %d hashes/-only records are not counted\n", r.HashOnly)
	}
	return writeTSV(out, r)
}
