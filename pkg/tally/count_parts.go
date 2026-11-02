//ff:func feature=count type=counter control=iteration dimension=1
//ff:what countParts — Layout이 나눈 조각 중 내장 레코드를 세고, 규격대로 나뉘지 않은 조각을 알린다
package tally

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// genericPayload is the name Layout gives a payload its spec did not fit.
const genericPayload = "payload"

// countParts counts the Record parts of a layout. A generic payload part
// means the record does not follow its spec, so it counts as unopened.
func countParts(parts []inspect.Part, spec codecspec.Spec, c *Counts) uint64 {
	var unopened uint64
	for _, p := range parts {
		if p.Name == genericPayload && !p.Record {
			unopened++
		}
		if !p.Record {
			continue
		}
		r, _, err := core.DecodeRecord(p.Words)
		if err != nil {
			unopened++
			continue
		}
		unopened += countRecord(r, p.Words, spec, c)
	}
	return unopened
}
