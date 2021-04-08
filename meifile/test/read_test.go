package meifile_test

import (
	"testing"

	"github.com/andytyc/goutils/meifile"
	"github.com/astaxie/beego/logs"
)

func TestRead01(t *testing.T) {
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

	err = meifile.WriteFileAppend(file1, []byte("\n==>让我们嗨起来吧~"))
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("WriteFileAppend ok", file1)

	data, err := meifile.ReadFile(file1)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("ReadFile ok", string(data))

	dir1Sub1 := dir1 + "/Subdir/"
	err = meifile.CreateDir(dir1Sub1)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("CreateDir ok", dir1Sub1)

	file2 := dir1 + "/" + "w2.log"
	err = meifile.CreateFile(file2)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("CreateFile ok", file2)

	fileList, err := meifile.ReadDir(dir1)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("ReadDir ok", fileList)
	for index, f := range fileList {
		logs.Info("r-f", index, f.Name(), f.IsDir(), f.ModTime(), f.Size(), f.Mode())
	}

	dirnameList, err := meifile.ReadScanDir(dir1)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("ReadScanDir ok", dirnameList)
}
