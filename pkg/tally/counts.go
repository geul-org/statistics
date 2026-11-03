//ff:type feature=count type=model
//ff:what Counts — 레코드 수 표 하나: TAG 9비트 코드(섹션 KKKK 4비트와 C 5비트)의 칸마다 한 개
package tally

// Counts holds one count per slot (slotOf): the 512 codes of the TAG space.
type Counts [512]uint64
