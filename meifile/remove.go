package meifile

import "os"

// Remove 删除 "文件" 或 "目录"及包含的目录/文件 {若路径已不存在,则不会报错}
func Remove(filepath string) (err error) {
	err = os.RemoveAll(filepath)
	if err != nil {
		return
	}
	return
}
