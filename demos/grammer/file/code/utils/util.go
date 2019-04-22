package utils

import (	
	"io/ioutil"
	"io"
	"bufio"
	"fmt"
	"log"
	"os"
)
/* 文件读写操作——file 按字节读取*/
func FileDemo()  {
	filePath := "file.txt"
	//1、打开文件
	// file, err := os.Open("file.txt") // For read access.	
	file, err := os.OpenFile(filePath,  os.O_RDWR | os.O_APPEND | os. O_CREATE,066) // For read access.	
	//2、关闭文件
	defer file.Close() 
	CheckError(err)
	// fmt.Println(file)

	// 3、写入数据 
	count, err := file.WriteString("welcomelearngo")
	CheckError(err)
	fmt.Printf("WriteString %d \n", count)
    //把string转成byte??
	str := "你好，世界！"
	data := []byte(str)
	count, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Write %d \n", count,)


	//4、读取文件数据 ???
	WriteFileByByte(filePath)
	//5.获取文件信息
	// GetFileInfo(file)
	
	
}

func WriteFileByByte(filePath string)  {
	// filePath = "bufio.txt"
	file, err := os.OpenFile(filePath,  os.O_RDWR, 066) 	
	data := make([]byte, 1024)//文件的信息可以读取进一个[]byte切片
	count, err := file.Read(data)	
	if err != nil {
		if err == io.EOF {//文件结尾
			fmt.Println("文件读取结束")	
		}
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}

/* 文件读写操作——buifo reader*/
//读取文件，带缓冲方式 按行读取数据
func BufioFileDemo()  {
	filePath := "bufio.txt"
	//1、打开文件	
	file, err := os.OpenFile(filePath,  os.O_RDWR | os.O_APPEND | os. O_CREATE,066) // For read access.	
	//2、关闭文件
	defer file.Close() 
	CheckError(err)
	// fmt.Println(file)

	// 3、写入数据 
	count, err := file.WriteString("welcomelearngo")
	CheckError(err)
	fmt.Printf("WriteString %d \n", count)
	
	 //4、读取数据 
	 WriteFileByLine(filePath)
}

func WriteFileByLine(filePath string)  {
	//1.打开文件获取到文件的指针(句柄)
	file , err := os.OpenFile(filePath,  os.O_RDWR, 066) 	
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	//使用defer关闭文件（延时）
	defer file.Close()
	//2.获取一个Reader(带缓冲), 通过 file 去构建reader
	reader := bufio.NewReader(file)

	//3.使用reader 读取文件, 循环的读取文件的内容
	for {
		con, err := reader.ReadString('\n') //读取一行
		fmt.Print(con)
		//如何判断文件读取完毕
		if err == io.EOF {
			fmt.Println("文件读取结束")
			break
		}
		
	}
}

/* 文件读写操作——iouti*/
//读取全部文件
func IoutiFileDemo()  {
	filePath := "iouti.txt"
	//1、打开文件	
	file, err := os.OpenFile(filePath,  os.O_RDWR | os.O_APPEND | os. O_CREATE,066) // For read access.	
	//2、关闭文件
	defer file.Close() 
	CheckError(err)
	// 3、写入数据 
	// data := []byte{115, 111, 109, 101, 10}
	str := "你好，世界！"
	data := []byte(str)
	
	 err = ioutil.WriteFile(filePath,data,066)
	CheckError(err)
	
	// filePath = "bufio.txt"
	// filePath = "file.txt"
	//4、读取文件
	res, err := ioutil.ReadFile(filePath)
	CheckError(err)
	fmt.Println(string(res))

}

/*文件拷贝*/
func CopyFileDemo(){
	w, err:= CopyFile("filecopy.txt","file.txt")
	CheckError(err)
	fmt.Println(w)
	
}

//文件统计
func CountDemo()  {
	filePath := "bufio.txt"
	CharCountImp(filePath)
}




