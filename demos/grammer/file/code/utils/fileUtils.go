package utils

import (
	"bufio"
	"log"
	"io"
	"fmt"
	"os"
)

//封装的文件拷贝方法
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

//判断文件或者目录存在的方法
func PathExists(path string) (bool, error) {
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
func GetFileInfo(file *os.File){
	//获取文件信息
	FileInfo, err :=file.Stat()
	CheckError(err)
	fmt.Println("FileInfo=",FileInfo)
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
	FileInfo.Name(),FileInfo.Size(),FileInfo.Mode(),FileInfo.ModTime(),FileInfo.IsDir(),FileInfo.Sys())
	
}


type charCount struct {
	ChCount    int //英文
	NumCount   int //数字 
	SpaceCount int //空格
	OtherCount int //其它字符数量
}
//统计英文、数字、空格和其他字符数量
func CharCountImp(filePath string) {
	file, err := os.Open(filePath) //说明这里需要事先准备一个 c:/test.txt 文件来进行测试
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	defer file.Close()

	var count charCount

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read file failed, err:%v", err)
			break
		}

		runeArr := []rune(str) //转成 []rune,可以处理中文字符
		for _, v := range runeArr {
		
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}

	}

	fmt.Printf("char count:%d\n", count.ChCount)
	fmt.Printf("num count:%d\n", count.NumCount)
	fmt.Printf("space count:%d\n", count.SpaceCount)
	fmt.Printf("other count:%d\n", count.OtherCount)
}
