package meitest

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
	logs.Info("simList  :", simList)
}
