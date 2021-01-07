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

// ReadScanDir 打开目录并扫描，返回其下一级"子目录名称"列表，按照文件名称{字符串的大小排序}进行排序
func ReadScanDir(path string) (list []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	list, err = f.Readdirnames(-1)
	if err != nil {
		return
	}
	sort.Slice(list, func(i, j int) bool { return list[i] < list[j] })
	return
}

// ReadFile 读取文件数据 |若想读取结果为字符串，则自己转换：string(data)
func ReadFile(path string) (data []byte, err error) {
	data, err = ioutil.ReadFile(path)
	if err != nil {
		// 若需要查看文件路径是否不存在？对err判定 => if err != nil && os.IsNotExist(err)
		return
	}
	return
}
