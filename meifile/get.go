package meifile

import "os"

// GetOpenFile 打开文件,获取可操作文件实例f {不建议使用(可参考如何打开文件)，若使用注意：f.Close()}
func GetOpenFile(path string, flag ...int) (*os.File, error) {
	var openflag int
	if flag == nil || len(flag) == 0 {
		openflag = os.O_RDWR | os.O_CREATE // 默认
	} else {
		openflag = flag[0]
	}
	// perm常用: os.ModePerm 或 0666
	f, err := os.OpenFile(path, openflag, os.ModePerm) // perm为0开头的0754中,754为linux下的文件权限
	if err != nil {
		return nil, err
	}
	// err == nil下,注意f使用完后,要执行f.Close(),否则程序长时间运行会挂掉,原因是内存溢出
	return f, nil
}
