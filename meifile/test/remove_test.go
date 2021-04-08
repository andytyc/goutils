package meifile_test

import (
	"testing"
	"time"

	"github.com/andytyc/goutils/meifile"

	"github.com/astaxie/beego/logs"
)

func TestRemove01(t *testing.T) {
	dir1 := "./logs/write"
	err := meifile.CreateDir(dir1)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("CreateDir ok", dir1)

	file1 := dir1 + "/" + "w1.log"
	err = meifile.CreateFile(file1)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("CreateFile ok", file1)

	file2 := dir1 + "/" + "w2.log"
	err = meifile.CreateFile(file2)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("CreateFile ok", file2)

	time.Sleep(4 * time.Second) // 4s时间，可以去查看下 目录/文件 已创建成功，后续测试删除操作

	err = meifile.Remove(file1)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("Remove ok", file1)

	err = meifile.Remove(file2)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("Remove ok", file2)

	err = meifile.Remove(dir1)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("Remove ok", dir1)
}
