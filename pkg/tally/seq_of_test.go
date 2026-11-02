//ff:func feature=live type=helper control=sequence
//ff:what TestSeqOf — pack 위치가 단계보다 앞선다
package tally

import "testing"

func TestSeqOf(t *testing.T) {
	if seqOf(1, 0) <= seqOf(0, 1<<32-1) || seqOf(2, 3) != 2<<32|3 {
		t.Fatal("order")
	}
}
