package meifile

import (
	"io"
	"os"
)

// writeFileData 写入内容至文件 {文件路径不存在则创建，然后写入数据}
//
// data传入类型：
// 写入字符串 => []byte(data)
// 写入二进制 => data
//
// perm推荐：
// os.ModePerm 或 0666
func writeFileData(path string, data []byte, flag int, perm os.FileMode) (err error) {
	// 参考:err = ioutil.WriteFile(filepath, data, os.ModePerm)
	f, err := os.OpenFile(path, flag, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WriteFile 写入内容至文件{不存在则创建文件} | 若传入字符串，则自己转换：[]byte(data)
func WriteFile(filepath string, data []byte) error {
	return writeFileData(filepath, data, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
}

// WriteFileAppend 追加内容到文件末尾 {不存在则创建文件}
func WriteFileAppend(path string, data []byte) error {
	return writeFileData(path, data, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
}
