package meifile_test

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
