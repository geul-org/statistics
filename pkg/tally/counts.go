//ff:type feature=count type=model
//ff:what Counts — 레코드 수 표 하나: Kind 2비트와 Codec 7비트의 칸마다 한 개
package tally

// Counts holds one count per slot (slotOf).
type Counts [512]uint64
