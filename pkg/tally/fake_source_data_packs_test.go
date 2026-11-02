//ff:func feature=archive type=helper control=sequence
//ff:what fakeSource.DataPacks — 시험용 pack 목록
package tally

// DataPacks returns the fixed list.
func (s fakeSource) DataPacks() []uint32 {
	return s.packs
}
