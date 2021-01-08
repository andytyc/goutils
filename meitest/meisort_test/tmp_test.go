package meisort_test

import (
	"sort"
	"testing"

	"github.com/andytyc/goutils/meisort"
	"github.com/astaxie/beego/logs"
)

func TestInt(t *testing.T) {
	// ------------------------------------------------------------------------------------
	// []Float32
	// ------------------------------------------------------------------------------------

	sims := []float32{0.2, 0.01, 0.4, 0.8, 0.6, 0.9, 0.99, 0.91, 0.0, 0.77}
	logs.Info("sims  :", sims)

	simList := meisort.SortIntFloat32s{}
	for index, sim := range sims {
		simList = append(simList, meisort.SortIntFloat32{Name: index, Value: sim})
	}
	sort.Sort(simList)
	logs.Info("simList  :", simList) // 这种方式，除了排序外，还能为每个值绑定一个键值{比如：字典的key,切片的index}
}

func TestSliceOO(t *testing.T) {
	sims := []float32{0.2, 0.01, 0.4, 0.8, 0.01, 0.6, 0.9, 0.99, 0.91, 0.0, 0.77}
	logs.Info("sims  :", sims)

	// 若仅仅是切片 {Slice}, 则可以：
	logs.Info("sims  :", sims)
	// sort.Slice(sims, func(i, j int) bool { return sims[i] < sims[j] }) // 升序
	sort.Slice(sims, func(i, j int) bool { return sims[i] > sims[j] }) // 降序

	// meisort.SliceSort(sims)

	logs.Info("sims  :", sims)
}
