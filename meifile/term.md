

# 介绍

术语中英文。



# os

## Flag

```bash
# wronly: 只写
os.O_WRONLY
# create: 创建
os.O_CREATE
# trunc: 截断
os.O_TRUNC
# append: 追加
os.O_APPEND
```


# FileMode

```go
// 默认内置自带
os.ModePerm FileMode = 0777 // Unix permission bits

// 自己也可以自定义，如：linux下的权限754，在golang格式转换示例：0xxx
var mode1 os.FileMode = 0754
```
