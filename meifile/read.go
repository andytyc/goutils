package meifile

import (
	"io/ioutil"
	"os"
	"sort"
)

// ReadDir 读取路径目录下包含的目录{紧邻目录}和文件列表
func ReadDir(path string) (files []os.FileInfo, err error) {
	files, err = ioutil.ReadDir(path)
	if err != nil {
		return
	}
	return
}

// ReadScanDir 打开目录并扫描，返回其下一级子目录名称列表，按照文件名称大小写进行排序
func ReadScanDir(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()
	list, err := f.Readdirnames(-1)
	if err != nil {
		return nil
	}
	sort.Slice(list, func(i, j int) bool { return list[i] < list[j] })
	return list
}

// ReadFile 读取文件数据
func ReadFile(path string) (data []byte, err error) {
	data, err = ioutil.ReadFile(path)
	if err != nil {
		// 若需要查看文件路径是否不存在？对err判定 => if err != nil && os.IsNotExist(err)
		return
	}
	// 若读取字符串，则：string(data)
	return
}
