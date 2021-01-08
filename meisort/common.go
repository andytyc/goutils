package meisort

import (
	"reflect"
	"sort"
)

// type SortBy []Type

// func (a SortBy) Len() int           { return len(a) }
// func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a SortBy) Less(i, j int) bool { return a[i] < a[j] } // 升序
// func (a SortBy) Less(i, j int) bool { return a[i] > a[j] } // 降序

// 重复性较多，使用代码生成器辅助

/*
1. 平时若只对 切片 进行排序，大可以直接调用方法。
=> sort.Slice()
2. 若，有些复杂逻辑排序，需要知道每个排序值的其他绑定属性值，比如：字典的key，切片的index等。
=> 自定义sort => 见：int.go, string.go...
*/

// SliceAsc 切片升序
func SliceAsc(slice interface{}) {
	sliceVal := []int{1, 2}
	sort.Slice(sliceVal, func(i, j int) bool { return sliceVal[i] < sliceVal[j] }) // 升序
}

// SliceDesc 切片降序
func SliceDesc() {

}

// LessDo LessDo
func LessDo(slice interface{}) func(i, j int) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic(&reflect.ValueError{Method: "LessDo", Kind: v.Kind()})
	}
	// Fast path for slices of size 0 and 1. Nothing to less.
	switch v.Len() {
	case 0:
		return func(i, j int) { panic("reflect: slice index out of range") }
	case 1:
		return func(i, j int) {
			if i != 0 || j != 0 {
				panic("reflect: slice index out of range")
			}
		}
	}

	// typ := v.Type().Elem().(*reflect.rtype)
	// size := typ.Size()
	// hasPtr := typ.ptrdata != 0

	// // Some common & small cases, without using memmove:
	// if hasPtr {
	// 	if size == reflectlite.ptrSize {
	// 		ps := *(*[]unsafe.Pointer)(v.ptr)
	// 		return func(i, j int) { ps[i], ps[j] = ps[j], ps[i] }
	// 	}
	// 	if typ.Kind() == reflectlite.String {
	// 		ss := *(*[]string)(v.ptr)
	// 		return func(i, j int) { ss[i], ss[j] = ss[j], ss[i] }
	// 	}
	// } else {
	// 	switch size {
	// 	case 8:
	// 		is := *(*[]int64)(v.ptr)
	// 		return func(i, j int) { is[i], is[j] = is[j], is[i] }
	// 	case 4:
	// 		is := *(*[]int32)(v.ptr)
	// 		return func(i, j int) { is[i], is[j] = is[j], is[i] }
	// 	case 2:
	// 		is := *(*[]int16)(v.ptr)
	// 		return func(i, j int) { is[i], is[j] = is[j], is[i] }
	// 	case 1:
	// 		is := *(*[]int8)(v.ptr)
	// 		return func(i, j int) { is[i], is[j] = is[j], is[i] }
	// 	}
	// }

	is := []int{}
	return func(i, j int) { is[i], is[j] = is[j], is[i] }
}

// // SliceSort SliceSort
// func SliceSort(items interface{}) {
// 	less := func(i, j int) bool { return items[i] > items[j] }
// 	sort.Slice(items, less) // 降序
// 	sort.Sort()
// }
