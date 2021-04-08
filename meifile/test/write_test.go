package meifile_test

import (
	"testing"

	"github.com/andytyc/goutils/meifile"
	"github.com/astaxie/beego/logs"
)

func TestWrite01(t *testing.T) {
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

	err = meifile.WriteFile(file1, []byte("你好 我的朋友~"))
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("WriteFile ok", file1)

	err = meifile.WriteFileAppend(file1, []byte("还好 好久不见~"))
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("WriteFileAppend ok", file1)
}
