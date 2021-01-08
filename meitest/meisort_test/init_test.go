package meisort_test

import (
	"github.com/astaxie/beego/logs"
)

func init() {
	logs.SetLogFuncCall(true)
	logs.SetLogFuncCallDepth(3)
	logs.SetLevel(logs.LevelInfo)
	logs.Async() // 异步跟踪，避免日志执行时间的计入
	logs.SetLogger(logs.AdapterConsole)
}

/*
// 测试执行到 Fatal 会打断，不会继续执行
t.Fatal(err)

// 对于t.log()进行日志记录，不如beego-logs清晰，日志也更好体现，这里初始化一个logs，方便使用
t.log("哈哈哈")
*/
