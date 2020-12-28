package meisort

// ------------------------------------------------------------------------------------
// 以Int作为key, 以其他可比较类型作为值，实现排序
// ------------------------------------------------------------------------------------

// Float32

// SortIntFloat32 SortIntFloat32
type SortIntFloat32 struct {
	Name  int
	Value float32
}

// SortIntFloat32s 通过对Value排序实现了sort.Interface接口
type SortIntFloat32s []SortIntFloat32

// Len Len
func (s SortIntFloat32s) Len() int {
	return len(s)
}

// Less Less
func (s SortIntFloat32s) Less(i, j int) bool {
	return s[i].Value > s[j].Value // 降序
}

// Swap Swap
func (s SortIntFloat32s) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Float64

// SortIntFloat64 SortIntFloat64
type SortIntFloat64 struct {
	Name  int
	Value float64
}

// SortIntFloat64s 通过对Value排序实现了sort.Interface接口
type SortIntFloat64s []SortIntFloat64

// Len Len
func (s SortIntFloat64s) Len() int {
	return len(s)
}

// Less Less
func (s SortIntFloat64s) Less(i, j int) bool {
	return s[i].Value > s[j].Value // 降序
}

// Swap Swap
func (s SortIntFloat64s) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
