//ff:func feature=cli type=command control=sequence
//ff:what run — 인자를 읽어 한도를 프로세스에 걸고 집계를 돌린다, 종료 코드를 돌려준다
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"runtime"
	"runtime/debug"

	"github.com/geul-org/geulstatic/pkg/tally"
)

// run sets the memory limit the caller chose for the whole process and counts
// the archives named by args. Threads start at 1: gogeul's zstd decoder for
// inventories takes its size from GOMAXPROCS at first use and every decoder
// keeps the last inventory it read, 4 GB in all for 28 on Wikidata. Inventories
// are read one at a time, so one is enough; tally.CountSet raises GOMAXPROCS
// for the documents. It returns 2 for a usage error and 1 when counting fails.
func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("geulstatic", flag.ContinueOnError)
	fs.SetOutput(stderr)
	memory := fs.Int64("memory", 0, "memory limit in bytes, half of it for path tracking; 0 for no limit")
	threads := fs.Int("threads", runtime.NumCPU(), "threads to use")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if fs.NArg() == 0 || *memory < 0 || *threads < 1 {
		fmt.Fprintln(stderr, "usage: geulstatic [-memory bytes] [-threads n] archive...")
		return 2
	}
	runtime.GOMAXPROCS(1)
	if *memory > 0 {
		debug.SetMemoryLimit(*memory)
	}
	err := tally.Run(context.Background(), fs.Args(), tally.Options{Memory: *memory, Threads: *threads}, stdout, stderr)
	if err != nil {
		fmt.Fprintln(stderr, "geulstatic:", err)
		return 1
	}
	return 0
}
