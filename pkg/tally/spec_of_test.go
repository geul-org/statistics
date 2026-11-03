//ff:func feature=archive type=resolver control=sequence
//ff:what TestSpecOf — 아카이브의 규격을 주소로 찾는다, 규격이 없으면 오류, 깨진 규격 문서는 오류
package tally

import (
	"context"
	"errors"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/pack"
)

func TestSpecOf(t *testing.T) {
	ctx := context.Background()
	wk := specArchive(t)
	empty := filepath.Join(t.TempDir(), "0x00000001")
	writeGen(t, empty, map[string][]byte{"0x00000001/docs/0x0000000000000001/0x00000001.geul": docOf(t)})
	archives, err := openArchives(ctx, []string{empty, wk})
	if err != nil {
		t.Fatal(err)
	}
	s, err := specOf(archives)
	if err != nil {
		t.Fatal(err)
	}
	if fs, ok := s[core.TargetWord(core.KindReference, core.CodecWikidataQID)]; !ok || len(fs) != 1 || len(s) != 1 {
		t.Fatalf("got %v", s)
	}
	if _, err := specOf(archives[:1]); !errors.Is(err, errNoSpec) {
		t.Fatalf("got %v", err)
	}
	broken := filepath.Join(t.TempDir(), "0x00000000")
	list := "0x00000000/codecs/0x00000140/versions.geul" // WIKIDATA-QID 0x014L
	writeGen(t, broken, map[string][]byte{list: docOf(t, core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("not a list")))})
	bad, err := pack.OpenDir(ctx, broken, pack.BuildOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := specOf([]*pack.Archive{bad}); err == nil {
		t.Fatal("broken list: want error")
	}
}
