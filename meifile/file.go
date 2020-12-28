package meifile

import (
	"fmt"
	"io/ioutil"
	"os"
)

// FileWrite 写入文件{不存在则创建文件} => []byte
func FileWrite(filepath string, data []byte) (err error) {
	err = ioutil.WriteFile(filepath, data, os.ModePerm)
	if err != nil {
		return
	}
	return
}

// FileWriteString 写入文件{不存在则创建文件} => string
func FileWriteString(filepath, data string) (err error) {
	err = ioutil.WriteFile(filepath, []byte(data), os.ModePerm)
	if err != nil {
		return
	}
	return
}

// FileRead 读取文件
func FileRead(filepath, msg string) (notexist bool, data []byte, err error) {
	data, err = ioutil.ReadFile(filepath)
	if err != nil && os.IsNotExist(err) {
		// 文件不存在
		notexist = true
		err = fmt.Errorf("read filepath error is not exist :%s", err)
	} else {
		// 报错异常
		err = fmt.Errorf("read filepath error :%s", err)
	}
	return
}
