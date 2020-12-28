package meitest

import (
	"github.com/astaxie/beego/logs"
)

func init() {
	logs.SetLogFuncCall(true)
	logs.SetLogFuncCallDepth(3)
	logs.SetLevel(logs.LevelInfo)
	logs.Async()
	logs.SetLogger(logs.AdapterConsole)
}

/*
go test -v -count=1 -run ^TestInt$ github.com/andytyc/goutils/meitest
go test -v -count=1 -timeout 30s -run ^TestInt$ github.com/andytyc/goutils/meitest
*/
