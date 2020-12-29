package meifile

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// PathRunPath 当前{程序执行}文件的绝对路径
// 比如测试运行的临时路径:/var/folders/ds/v2_y1w6s2c74ldx9gzv9kfzw0000gq/T/go-build468352716/b001/meitest.test
func PathRunPath() string {
	p, _ := filepath.Abs(os.Args[0])
	return p
}

// PathRunDir 当前{程序执行}文件的所在目录绝对路径
// 比如测试运行的临时路径:/var/folders/ds/v2_y1w6s2c74ldx9gzv9kfzw0000gq/T/go-build468352716/b001
func PathRunDir() string {
	return filepath.Dir(PathRunPath())
}

// PathDir 路径{文件/目录}的所在目录
// 备注：去除路径/的最后一部分(包含最后的/也去除 - 根路径/除外)
/*
可以解析:
	../meitest/logs/ => ../meitest/logs
	../meitest/logs => ../meitest
	../meitest/logs/22.log => ../meitest/logs
	../ => ..
	./ => .
	/ => /
	/home => /
不可以解析(空字符串 或 无/分隔符), 返回.:
	.. => .
	. => .
	"" => .
应用：
	* 所在目录:
		PathDir(path)
	* 上级目录:
		PathDir(PathDir(path))
*/
func PathDir(path string) string {
	return filepath.Dir(path)
}

// PathParentDir 路径的所在目录的上级目录
func PathParentDir(path string) string {
	return filepath.Dir(filepath.Dir(path))
}

// PathInfo 获取{文件或目录}信息 {ll path}
func PathInfo(path string) *os.FileInfo {
	info, err := os.Stat(path)
	if err != nil {
		return nil
	}
	return &info
}

// PathChmod 修改{文件/目录}权限 {chmod}
func PathChmod(path string, mode os.FileMode) error {
	return os.Chmod(path, mode)
}

// PathMove 文件移动/重命名 {mv src dst}
func PathMove(src string, dst string) error {
	return os.Rename(src, dst)
}

// PathFileCopy 文件复制 {cp = 只能是文件哦~}
func PathFileCopy(src string, dst string) error {
	srcf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcf.Close()
	dstf, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstf.Close()
	_, err = io.Copy(dstf, srcf) // 复制: read from srcf => write to dstf (写入是：覆盖不是追加)
	if err != nil {
		return err
	}
	err = dstf.Sync() // 提交内存数据写入磁盘
	if err != nil {
		return err
	}
	return nil
}

// ------------------------ bool 判定 ------------------------

// PathIsExist 路径文件/文件夹是否存在
func PathIsExist(filepath string) bool {
	_, err := os.Stat(filepath) // os.Stat获取文件信息
	if err != nil && !os.IsExist(err) {
		return false
	} else {
		return true
	}
}

// PathIsDir 路径是否为文件夹
func PathIsDir(path string) bool {
	var exist bool
	f, err := os.Stat(path)
	if err != nil && !os.IsExist(err) {
		exist = false
	} else {
		exist = true
	}
	if exist {
		return f.IsDir()
	} else {
		return false
	}
}

// PathIsFile 路径是否为文件
func PathIsFile(path string) bool {
	var exist bool
	f, err := os.Stat(path)
	if err != nil && !os.IsExist(err) {
		exist = false
	} else {
		exist = true
	}
	if exist {
		return !f.IsDir()
	} else {
		return false
	}
}

// PathFileIsReadable 文件是否可读
func PathFileIsReadable(path string) (readable bool) {
	readable = true
	f, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		readable = false
	}
	defer f.Close()
	return
}

// ------------------------ curd 增删改 ------------------------

// PathAddPrefix 是否以xx开头{没有自动添加}
func PathAddPrefix(filepath string, prefix string) string {
	if filepath == "" || filepath == "/" {
		return filepath
	}
	if !strings.HasPrefix(filepath, prefix) {
		filepath = prefix + filepath
	}
	return filepath
}

// PathAddSuffix 是否以xx结尾{没有自动添加}
func PathAddSuffix(filepath string, suffix string) string {
	if filepath == "" || filepath == "/" {
		return filepath
	}
	if !strings.HasSuffix(filepath, suffix) {
		filepath += suffix
	}
	return filepath
}

// PathSubPrefix 是否以xx开头{有则删除}
func PathSubPrefix(filepath string, prefix string) string {
	if filepath == "" || filepath == "/" {
		return filepath
	}
	if strings.HasPrefix(filepath, prefix) {
		filepath = filepath[len(prefix):]
	}
	return filepath
}

// PathSubSuffix 是否以xx结尾{有则删除}
func PathSubSuffix(filepath string, suffix string) string {
	if filepath == "" || filepath == "/" {
		return filepath
	}
	if strings.HasSuffix(filepath, suffix) {
		pathlen := len(filepath)
		filepath = filepath[:pathlen-len(suffix)]
	}
	return filepath
}

// PathExt 路径的文件扩展名
//
// filepath.Ext(path) 执行示例：
// 22.log => .log
// ../meitest/logs/22.log => .log
// ../meitest/logs/22. => .
// ../meitest/logs/22 => ""
// ../meitest/logs/ => ""
// ../meitest/logs => ""
// / => ""
// "" => ""
//
// 更新：
// 在以上基础上，进行需求更新。=> {有,无}
func PathExt(path string) string {
	if path == "" {
		return ""
	}
	ext := filepath.Ext(path)
	if ext == "" || ext == "." {
		return ""
	}
	return ext
}

// PathBase 路径的文件名称
// 备注：没有/的就是全部字符串，有/的则是最后一个/后一部分元素
//
// filepath.Base(path) 执行示例：
// 22.log => 22.log
// ../meitest/logs/22.log => 22.log
// ../meitest/logs/22. => 22.
// ../meitest/logs/22 => 22
// ../meitest/logs/ => logs
// ../meitest/logs => logs
// / => /
// "" => .
//
// 更新：
// 在以上基础上，进行需求更新。=> {有,无}
func PathBase(path string) string {
	if path == "" {
		return ""
	}
	baseItem := filepath.Base(path)
	if baseItem == "." || baseItem == "/" {
		return ""
	}
	return baseItem
}

// PathJoin 将多个路径拼接为一个路径。常见应用：路径拼接 或 基础路径拼接相对路径
// filepath.Join() 将任意数量的路径元素连接到一条路径中，并在必要时添加一个分隔符。
// 加入通话清理结果；特别是，所有空字符串都将被忽略。
// 在Windows上，且仅当第一个路径元素是UNC路径时，结果才是UNC路径。
// ../meitest/logs | 24.log => ../meitest/logs/24.log
// ../meitest/logs | /24.log => ../meitest/logs/24.log
// ../meitest/logs | hh/24.log => ../meitest/logs/hh/24.log
// ../meitest/logs | ./hh/24.log => ../meitest/logs/hh/24.log
// ../meitest/logs | ../hh/24.log => ../meitest/hh/24.log
func PathJoin(elempath ...string) string {
	return filepath.Join(elempath...)
}

// PathRel 获取{targpath路径}相对于{basepath路径}的相对路径
// 也就是说，Join（basepath，Rel（basepath，targpath））等同于targpath本身。
// 成功后，即使basepath和targpath不共享任何元素，返回的路径也始终相对于basepath。
// 如果无法相对于基本路径创建targpath，或者如果需要知道当前工作目录以进行计算，则会返回错误。Rel在结果上调用Clean。
//
// 对于俩路径来说，计算时是都按照是{目录路径}才进行相对处理的。
// ../meitest/logs/sub | ../meitest/logs/sub => .
// ../meitest/logs/sub/ | ../meitest/logs/sub => .
// ../meitest/logs/sub | ../meitest/logs/sub/ => .
// ../meitest/logs/sub/ | ../meitest/logs => ..
// ../meitest/logs/sub/ | ../meitest/logs/sub2 => ../sub2
func PathRel(basepath, targpath string) (string, error) {
	return filepath.Rel(basepath, targpath)
}

// PathSplit 路径分割
// ../meitest/logs/sub2 => ../meitest/logs/ | sub2
// ../meitest/logs/sub2/ => ../meitest/logs/sub2/ | ""
// ../meitest/logs/sub2/22.log=> ../meitest/logs/sub2/ | 22.log
func PathSplit(path string) (dir string, file string) {
	return filepath.Split(path)
}
