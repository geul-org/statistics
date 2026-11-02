//ff:func feature=archive type=resolver control=iteration dimension=1
//ff:what specOf — 주어진 아카이브의 0x0000에서 규격 문서를 주소로 찾아 모은다, 아카이브를 나열하지 않는다
package tally

import (
	"errors"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/pack"
	"github.com/geul-org/gogeul/pkg/resolve"
)

// errNoSpec means no archive given holds the codec spec documents.
var errNoSpec = errors.New("no codec spec documents under 0x0000: also give the archive or set that holds them")

// specOf reads the specs by address (resolve.Dir.CodecSpecs): resolve.Dir.CodecSpec
// lists the Issuer through Archive.ReadDir, which loads every path of every set
// into memory. Earlier archives win. Without specs no embedded record would be
// counted, so that is an error.
func specOf(archives []*pack.Archive) (codecspec.Spec, error) {
	s := codecspec.Spec{}
	for _, a := range archives {
		o, err := resolve.Dir{Store: a}.CodecSpecs(specTargets())
		if err != nil {
			return nil, err
		}
		s = s.Merge(o)
	}
	if len(s) == 0 {
		return nil, errNoSpec
	}
	return s, nil
}
