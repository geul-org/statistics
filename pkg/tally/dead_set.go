//ff:type feature=live type=model
//ff:what deadSet — pack별 죽은 항목 비트, 인벤토리 항목 하나에 1비트
package tally

// deadSet has one bit per inventory step, by pack position: set when a later
// step names the same path, so the step's record is no longer that path's.
type deadSet [][]uint64
