package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// 打开文件操作
func OpenFile(fileName string) (file *os.File) {
	flag, err := PathIsExist(fileName)
	if flag {
		//如果文件存在
		file, err = os.OpenFile(fileName, os.O_APPEND, 0666)
	} else {
		//创建文件
		file, err = os.Create(fileName)
	}
	CheckError(err)
	return file
}

//判断文件或者目录存在的方法
func PathIsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { // 文件或者目录存在
		log.Fatal(err)
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//获取文件信息
func GetFileInfo(file *os.File) {
	//获取文件信息
	FileInfo, err := file.Stat()
	CheckError(err)
	fmt.Println("FileInfo=", FileInfo)
	/*
			type FileInfo interface {
		    Name() string       // 文件的名字（不含扩展名）
		    Size() int64        // 普通文件返回值表示其大小；其他文件的返回值含义各系统不同
		    Mode() FileMode     // 文件的模式位
		    ModTime() time.Time // 文件的修改时间
		    IsDir() bool        // 等价于Mode().IsDir()
		    Sys() interface{}   // 底层数据来源（可以返回nil）
		    }
	*/
	fmt.Printf("Name=%v \nSize=%d \nMode= %v \nModTime=%v \nIsDir=%v \nSys=%v \n",
		FileInfo.Name(), FileInfo.Size(), FileInfo.Mode(), FileInfo.ModTime(), FileInfo.IsDir(), FileInfo.Sys())

}

//文件拷贝方法
func CopyFile(dstName string, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	CheckError(err)
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	CheckError(err)
	defer dst.Close()
	//拷贝文件
	return io.Copy(dst, src)
}

// 写入文件
func WriteToFile(file *os.File, content string) {
	writer := bufio.NewWriter(file)
	writer.WriteString(content)
	writer.WriteString("\r\n")
	writer.Flush()
}

func ReadFormFile() {

}
