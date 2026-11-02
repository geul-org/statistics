//ff:func feature=report type=helper control=sequence
//ff:what failWriter.Write — 언제나 실패한다
package tally

import "errors"

// Write fails.
func (failWriter) Write(p []byte) (int, error) {
	return 0, errors.New("write failed")
}
