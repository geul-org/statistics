//ff:func feature=count type=helper control=sequence
//ff:what TestFirstError — 원인이 있으면 원인, 없으면 인벤토리 오류
package tally

import (
	"errors"
	"testing"
)

func TestFirstError(t *testing.T) {
	a, b := errors.New("a"), errors.New("b")
	if firstError(a, b) != a || firstError(nil, b) != b || firstError(nil, nil) != nil {
		t.Fatal("pick")
	}
}
