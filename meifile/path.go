package meifile

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

// PathCheckPrefix 是否以xx开头{没有自动添加}
func PathCheckPrefix(filepath string, prefix string) string {
	if filepath == "" {
		return filepath
	}
	if !strings.HasPrefix(filepath, prefix) {
		filepath = prefix + filepath
	}
	return filepath
}

// PathCheckSuffix 是否以xx结尾{没有自动添加}
func PathCheckSuffix(filepath string, suffix string) string {
	if filepath == "" {
		return filepath
	}
	if !strings.HasSuffix(filepath, suffix) {
		filepath += suffix
	}
	return filepath
}

// PathGetLastItem 获取路径的最后一个 "/" 的后面部分
func PathGetLastItem(filepath string) string {
	if filepath == "" {
		return ""
	}
	lastItem := path.Base(filepath)
	if lastItem == "." || lastItem == "/" {
		return ""
	}
	return lastItem
}

// PathNameGetExt 获取文件名{如：xxx.png}的后缀, ""表示无后缀
func PathNameGetExt(filename string) string {
	if filename == "" {
		return ""
	}
	ext := path.Ext(filename)
	if ext == "." {
		return ""
	}
	return ext
}

// PathGetExt 获取路径文件后缀 xx/xx/xx.jpg => .jpg
func PathGetExt(filepath string) string {
	return PathNameGetExt(PathGetLastItem(filepath))
}

// PathExists 判断所给路径文件/文件夹是否存在
func PathExists(filepath string) bool {
	_, err := os.Stat(filepath) // os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// PathIsDir 判断所给路径是否为文件夹
func PathIsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// PathIsFile 判断所给路径是否为文件
func PathIsFile(path string) bool {
	return !PathIsDir(path)
}

// PathRemoveAll 删除 "文件" 或 "目录"及包含的目录/文件 {若路径已不存在,则不会报错}
func PathRemoveAll(filepath, msg string) (err error) {
	err = os.RemoveAll(filepath)
	if err != nil {
		return
	}
	return
}

// PathFileCreate 创建文件
func PathFileCreate(path string) error {
	dir := Dir(path)
	if !PathExists(dir) {
		DirExistOrCreateAll(dir)
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}

// 获取当前执行文件的绝对路径
func SelfPath() string {
	p, _ := filepath.Abs(os.Args[0])
	return p
}

// 获取当前执行文件的目录绝对路径
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

// 获取指定文件路径的文件名称
func Basename(path string) string {
	return filepath.Base(path)
}

// 获取指定文件路径的目录地址绝对路径
func Dir(path string) string {
	return filepath.Dir(path)
}

// 获取指定文件路径的文件扩展名
func Ext(path string) string {
	return filepath.Ext(path)
}
