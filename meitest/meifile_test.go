package meitest

import (
	"testing"

	"github.com/andytyc/goutils/meifile"
	"github.com/astaxie/beego/logs"
)

func TestPath(t *testing.T) {
	// filepath1 := ""
	// filepath1 := "/"
	// filepath1 := "/home/andytyc"
	// filepath1 := "./logs"

	// filepath1 := "../meitest/logs"
	// filepath1 := "../meitest/logs/"
	filepath1 := "../meitest/logs/22.log"
	// filepath1 := "../meitest/logs/22."
	// filepath1 := "../meitest/logs/22"
	// filepath1 := "22.log"

	// filepath1 := "../meitest/logs-2"
	// filepath1 := "../meitest/logs-2/22.log"

	// path
	exist := meifile.PathIsExist(filepath1)
	logs.Info("exist :", exist)

	isdir := meifile.PathIsDir(filepath1)
	logs.Info("isdir :", isdir)

	isfile := meifile.PathIsFile(filepath1)
	logs.Info("isfile :", isfile)

	curpath := meifile.PathRunPath()
	logs.Info("curpath :", curpath)

	curdir := meifile.PathRunDir()
	logs.Info("curdir :", curdir)

	// dir
	dir := meifile.PathDir(filepath1)
	logs.Info("dir :", dir)

	pdir := meifile.PathParentDir(filepath1)
	logs.Info("pdir :", pdir)

	ext := meifile.PathExt(filepath1)
	logs.Info("ext :", ext)

	name := meifile.PathBase(filepath1)
	logs.Info("name :", name)

	// filepath2 := "24.log"
	// filepath2 := "/24.log"
	// filepath2 := "hh/24.log"
	// filepath2 := "./hh/24.log"
	filepath2 := "../hh/24.log"
	// filepath2 := "./logs/22.log"

	filepath3 := "../meitest/logs/sub"

	// filepath4 := "../meitest/logs/22.log"
	// filepath4 := "../meitest/logs"
	// filepath4 := "../meitest/logs/"
	filepath4 := "../meitest/logs/sub2"

	fpath1 := meifile.PathJoin(filepath1, filepath2)
	logs.Info("fpath1 :", fpath1)

	fpath2, err := meifile.PathRel(filepath3, filepath4)
	if err != nil {
		t.Error("err :", err)
	}
	logs.Info("fpath2 :", fpath2)

	// filepath5 := "../meitest/logs/sub2"
	// filepath5 := "../meitest/logs/sub2/"
	filepath5 := "../meitest/logs/sub2/22.log"

	fpath3, fpath4 := meifile.PathSplit(filepath5)
	logs.Info("fpath3 :", fpath3, "fpath4 :", fpath4)

	// filepath6 := "../meitest/logs/sub2"
	// filepath6 := "/home/andytyc/meitest/logs/sub2"
	filepath6 := "/home/andytyc/meitest/logs/sub2/"

	// fpath5 := meifile.PathAddSuffix(filepath6, "/")
	// fpath5 := meifile.PathAddPrefix(filepath6, "/")
	fpath5 := meifile.PathSubSuffix(filepath6, "/")
	// fpath5 := meifile.PathSubPrefix(filepath6, "/")
	logs.Info("fpath5 :", fpath5)
}

func TestDir(t *testing.T) {
	// filepath1 := "../meitest/logs/sub222"
	// filepath1 := "../meitest/logs/sub333/"
	filepath1 := "../meitest/logs/sub2"

	err := meifile.CreateDir(filepath1)
	if err != nil {
		t.Error("err :", err)
	}
	logs.Info("create dir :", filepath1)

	filepath2 := "../meitest/logs"
	filelist := meifile.ReadScanDir(filepath2)
	logs.Info("filelist :", filelist)
}

func TestFile(t *testing.T) {
	// filepath1 := "../meitest/logs/sub2/22.log"
	filepath1 := "../meitest/logs/sub2/23.log"
	// filepath11 := "../meitest/logs/sub2/233.log"
	// filepath11 := "../meitest/logs/sub2-22/233.log"
	filepath11 := "../meitest/logs/sub2-22/233.log"

	err := meifile.CreateFile(filepath1)
	if err != nil {
		t.Error("err create :", err)
	}
	logs.Info("create file :", filepath1)

	filepath2 := "../meitest/logs/sub2/24.log"

	err = meifile.PathFileCopy(filepath1, filepath2)
	if err != nil {
		t.Error("err copy :", err)
	}
	logs.Info("copy file :", filepath2)

	err = meifile.PathMove(filepath1, filepath11)
	if err != nil {
		t.Error("err move/rename :", err)
	}
	logs.Info("move/rename file :", filepath11)
}
