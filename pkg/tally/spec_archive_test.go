//ff:func feature=archive type=helper control=sequence
//ff:what specArchive — 시험용: 0x00000000/codecs/에 코덱 문서 하나를 공표한 `.well-known/geul` pack 폴더
package tally

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/pack"
	"github.com/geul-org/gogeul/pkg/resolve"
)

// specArchive publishes the spec of QID references and packs it as Issuer sets.
func specArchive(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	target := core.TargetWord(core.KindReference, core.CodecWikidataQID)
	if _, err := (resolve.Dir{Root: root}).PublishCodec(codecDocOf(t, target)); err != nil {
		t.Fatal(err)
	}
	wk := filepath.Join(t.TempDir(), "geul")
	if _, err := pack.Split(context.Background(), root, wk, pack.BuildOptions{}); err != nil {
		t.Fatal(err)
	}
	return wk
}
