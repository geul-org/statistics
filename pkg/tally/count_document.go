//ff:func feature=count type=counter control=iteration dimension=1
//ff:what CountDocument — 문서 바이트의 레코드를 core 파서로 차례로 잘라 내장 레코드까지 센다
package tally

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
)

// CountDocument adds every record of the document b to c, the structure
// records DOC-*, BODY and END included, and returns how many records could
// not be opened (countRecord).
func CountDocument(b []byte, spec codecspec.Spec, c *Counts) (uint64, error) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return 0, err
	}
	var unopened uint64
	for i := 0; i < len(ws); {
		r, n, err := core.DecodeRecord(ws[i:])
		if err != nil {
			return 0, err
		}
		unopened += countRecord(r, ws[i:i+n], spec, c)
		i += n
	}
	return unopened, nil
}
