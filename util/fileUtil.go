package util

import (
	"io/fs"
	"io/ioutil"
	"os"
)

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// write file with name, data, perm
func WriteCover(fileName string, data []byte, perm uint32) {
	// 覆盖式写入
	DropErr(ioutil.WriteFile(fileName, data, fs.FileMode(perm))) // perm 0664 默认
}

// read file
func Read(fileName string) []byte {
	open, err := os.Open(fileName)
	all, err := ioutil.ReadAll(open)
	DropErr(err)
	return all
}

// return file obj
/**
其中，flag 有以下几种常用的值：

os.O_CREATE: create if none exists 不存在则创建
os.O_RDONLY: read-only 只读
os.O_WRONLY: write-only 只写
os.O_RDWR: read-write 可读可写
os.O_TRUNC: truncate when opened 文件长度截为0：即清空文件
os.O_APPEND: append 追加新数据到文件
*/
func OpenFileChannel(fileName string, flag int, perm uint32) *os.File {
	file, err := os.OpenFile(fileName, flag, fs.FileMode(perm))
	DropErr(err)
	return file
}

// 添加 html中的 下拉框的 前半部分 ，中间是数据写入
func WriteSelectItem1(page string) string {
	return `<details>
<summary>` + page + `</summary>`
}

// 添加 html中的下拉框的 后半部分，中间是数据写入
func WriteSelectItem2() string {
	return `</details>`
}

// 清空文件
func EmptyFile(fileName string) {
	DropErr(ioutil.WriteFile(fileName, []byte(""), 0664))
}
