//ff:func feature=cli type=command control=sequence
//ff:what main — geulstatic 명령 진입점
package main

import "os"

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
