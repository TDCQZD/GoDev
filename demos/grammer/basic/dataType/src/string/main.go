package main
import "fmt"
func main(){
	var name string ="zhang san" //字符串不可更改
	fmt.Printf("name[0]=%v\n",name[0])
	fmt.Printf("name[0]=%c\n",name[0])

	//双引号可以识别转义字符
	str := "asjfd a slj \nascxcas"
	fmt.Println(str)
    //反引号，原样输出
	str1 :=`package main
	// import "fmt"
	func main(){	
	}`
	fmt.Printf(str1)
}