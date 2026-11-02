//ff:type feature=archive type=model
//ff:what Options — 실행자가 정하는 메모리 한도와 스레드 수
package tally

// Options are the limits of one run. Memory is in bytes: half of it bounds
// the path tracking table and 0 means no limit. Threads is how many documents
// are worked on at once, at least 1.
type Options struct {
	Memory  int64
	Threads int
}
