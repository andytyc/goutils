package sortutil

// ------------------------------------------------------------------------------------
// Int64
// ------------------------------------------------------------------------------------

// SortStringInt64 SortStringInt64
type SortStringInt64 struct {
	Name  string
	Value int64
}

// SortStringInt64s 通过对Value排序实现了sort.Interface接口
type SortStringInt64s []SortStringInt64

// Len Len
func (s SortStringInt64s) Len() int {
	return len(s)
}

// Less Less
func (s SortStringInt64s) Less(i, j int) bool {
	return s[i].Value > s[j].Value // 降序排序
}

// Swap Swap
func (s SortStringInt64s) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
