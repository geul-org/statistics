//ff:func feature=live type=tracker control=sequence
//ff:what TestMarkOf — 같은 경로는 같은 해시, 다른 경로는 다른 해시, 순번을 지닌다
package tally

import "testing"

func TestMarkOf(t *testing.T) {
	a, b, c := markOf("x.geul", 1), markOf("x.geul", 2), markOf("y.geul", 1)
	if a.h0 != b.h0 || a.h1 != b.h1 || a.seq != 1 || b.seq != 2 {
		t.Fatalf("%+v %+v", a, b)
	}
	if a.h0 == c.h0 && a.h1 == c.h1 {
		t.Fatal("different paths share a hash")
	}
}
