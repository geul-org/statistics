//ff:type feature=report type=model
//ff:what failWriter — 시험용: 쓰기마다 실패하는 Writer
package tally

// failWriter fails every write.
type failWriter struct{}
