package meifile_test

import (
	"os"
	"testing"

	"github.com/andytyc/goutils/meifile"
	"github.com/astaxie/beego/logs"
)

func TestPath01(t *testing.T) {
	//********************************************************
	// 准备
	//********************************************************

	dir1 := "./logs/write"
	err := meifile.CreateDir(dir1)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("CreateDir ok", dir1)
	dir1Sub1 := dir1 + "/Subdir1/"
	err = meifile.CreateDir(dir1Sub1)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("CreateDir ok", dir1Sub1)

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

	file2 := dir1 + "/" + "w2.log"
	err = meifile.CreateFile(file2)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("CreateFile ok", file2)
	err = meifile.WriteFile(file2, []byte("\n还好 好久不见~"))
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("WriteFile ok", file2)

	logs.Info("********************************************************")

	//********************************************************
	// 路径判定
	//********************************************************

	isExist := meifile.PathIsExist(dir1)
	logs.Info("PathIsExist dir", isExist, dir1)
	isExist = meifile.PathIsExist(file1)
	logs.Info("PathIsExist file", isExist, file1)

	isDir := meifile.PathIsDir(dir1)
	logs.Info("PathIsDir", isDir, dir1)
	isDir = meifile.PathIsDir(file1)
	logs.Info("PathIsDir", isDir, file1)

	isFile := meifile.PathIsFile(dir1)
	logs.Info("PathIsFile", isFile, dir1)
	isFile = meifile.PathIsFile(file1)
	logs.Info("PathIsFile", isFile, file1)

	isRead := meifile.PathIsReadable(dir1)
	logs.Info("PathIsReadable", isRead, dir1)
	isRead = meifile.PathIsFile(file1)
	logs.Info("PathIsReadable", isRead, file1)

	logs.Info("********************************************************")

	//********************************************************
	// 路径裁剪
	//********************************************************

	dir1Sub2 := dir1 + "/Subdir2"
	dir1Sub2File := dir1 + "/Subdir2/22.log"

	tmp1 := meifile.PathDoAddPrefix(dir1Sub2, "/")
	logs.Info("PathDoAddPrefix", tmp1)
	tmp1 = meifile.PathDoDelPrefix(dir1Sub2, "/")
	logs.Info("PathDoDelPrefix", tmp1)

	tmp1 = meifile.PathDoAddSuffix(dir1Sub2, "/")
	logs.Info("PathDoAddSuffix", tmp1)
	tmp1 = meifile.PathDoDelSuffix(dir1Sub2, "/")
	logs.Info("PathDoDelSuffix", tmp1)

	tmp1 = meifile.PathDoGetDir(dir1Sub2)
	logs.Info("PathDoGetDir", tmp1)
	tmp1 = meifile.PathDoGetPDir(dir1Sub2)
	logs.Info("PathDoGetPDir", tmp1)

	tmp1 = meifile.PathDoGetBase(dir1Sub2File)
	logs.Info("PathDoGetBase", tmp1)
	tmp1 = meifile.PathDoGetExt(dir1Sub2File)
	logs.Info("PathDoGetExt", tmp1)

	tmp1 = meifile.PathDoJoin(dir1Sub2, dir1Sub2File)
	logs.Info("PathDoJoin", tmp1)

	tmp1, err = meifile.PathDoRel(dir1Sub2, dir1Sub2File)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("PathDoRel", tmp1)

	tmp1, tmp2 := meifile.PathDoSplit(dir1Sub2File)
	logs.Info("PathDoSplit", tmp1, tmp2)

	logs.Info("********************************************************")

	//********************************************************
	// 路径文件/目录实例 File
	//********************************************************

	tmp1 = meifile.PathFileRunPath()
	logs.Info("PathFileRunPath", tmp1)
	tmp1 = meifile.PathFileRunDir()
	logs.Info("PathFileRunDir", tmp1)

	tmp1f, err := meifile.PathFileGetInfo(tmp1)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("PathFileGetInfo", tmp1f.Name(), tmp1f.IsDir(), tmp1f.Size(), tmp1f.ModTime(), tmp1f.Mode())

	tocopyfile1 := dir1 + "/" + "copy1.log"
	err = meifile.PathFileCopy(file1, tocopyfile1)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("PathFileCopy ok", tocopyfile1)

	tocopyfile1f, err := meifile.PathFileGetInfo(tocopyfile1)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("PathFileGetInfo", tocopyfile1f.Name(), tocopyfile1f.IsDir(), tocopyfile1f.Mode())

	var mode1 os.FileMode = 0777
	err = meifile.PathFileChmod(tocopyfile1, mode1)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("PathFileChmod ok", mode1, tocopyfile1)

	tocopyfile1f, err = meifile.PathFileGetInfo(tocopyfile1)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("PathFileGetInfo", tocopyfile1f.Name(), tocopyfile1f.IsDir(), tocopyfile1f.Mode())

	tomovefile1 := dir1 + "/" + "move1.log"
	err = meifile.PathFileMove(tocopyfile1, tomovefile1)
	if err != nil {
		t.Fatal(err)
	}
	logs.Info("PathFileMove ok", tomovefile1)
}
