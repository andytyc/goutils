package meifile

import (
	"io/ioutil"
	"os"
)

// DirExistOrCreateAll 存在：不操作并返回nil; 不存在：则创建目录{一级/多级都可以}
func DirExistOrCreateAll(dirpath string) (err error) {
	err = os.MkdirAll(dirpath, os.ModePerm)
	if err != nil {
		return
	}
	return
}

// DirRead 读取路径目录下包含的目录{紧邻目录}和文件列表
func DirRead(dirpath string) (files []os.FileInfo, err error) {
	files, err = ioutil.ReadDir(dirpath)
	if err != nil {
		return
	}
	return
}
