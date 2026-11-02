//ff:func feature=archive type=helper control=sequence
//ff:what fakeSource.Record — 시험용 문서 읽기
package tally

import (
	"context"

	"github.com/geul-org/gogeul/pkg/pack"
)

// Record returns the document at l.Off, or err when recFail is set.
func (s fakeSource) Record(ctx context.Context, l pack.Loc) ([]byte, error) {
	if s.recFail {
		return nil, s.err
	}
	return s.docs[l.Off], nil
}
