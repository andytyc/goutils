package meifile

import (
	"os"
)

// CreateDir 存在：不操作并返回nil; 不存在：则创建目录{一级/多级都可以}
func CreateDir(dirpath string) (err error) {
	err = os.MkdirAll(dirpath, os.ModePerm)
	if err != nil {
		return
	}
	return
}

// CreateFile 存在：不操作并返回nil; 不存在：则创建文件
func CreateFile(filepath string) (err error) {
	exist := PathIsExist(filepath)
	if exist {
		return
	}
	dir := PathDir(filepath)
	if !PathIsExist(dir) {
		err = CreateDir(dir)
		if err != nil {
			return
		}
	}
	f, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer f.Close()
	return
}
