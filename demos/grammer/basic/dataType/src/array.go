package utils

import (
	"time"
	"math/rand"
	"fmt"
)


/*
一个养鸡场有6只鸡，它们的体重分别是3kg,5kg,1kg,
3.4kg,2kg,50kg 。请问这六只鸡的总体重是多少?平
均体重是多少? 请你编一个程序
*/
func ArrayDemo()  {
	var hens [6] float64
	hens[0]=3
	hens[1]=5
	hens[2]=1
	hens[3]=3.4
	hens[4]=2
	hens[5]=50
	var sum float64
	for _, value := range hens {
		sum +=value
	}

	aver := sum / float64(len(hens))

	fmt.Printf("这六只鸡的总体重是%f 平均体重是%f \n",sum,aver)
}

//数组的四种初始化方式
func ArrayDemo2()  {

	//定义形式
	var array [3] int = [3] int {1,2,3} 
	var array1  = [3] int {1,2,3}
	var array2  = [...] int {1,2,3}
	//指定元素对应下标
	var array3 = [3] int {1:1,2:2,0:3}
	fmt.Println("array=",array)
	fmt.Println("array1=",array1)
	fmt.Println("array2=",array2)
	fmt.Println("array3=",array3)
}

/* 请求出一个数组的最大值，并得到对应的下标
分析：
1、定义数组
2、循环遍历数组
3、比较 获取 max maxIndex
*/
func ArrayDemo3()  {

	
	var array =[...] int {12,3333,89,0,1,34222,664}
	max := array[0]
	maxIndex := 0
	for index ,value:= range array {
		if max < value {
			max = value
			maxIndex = index
		}
	}
   fmt.Printf("数组= %v 中 max= %d maxIndex= %d",array,max,maxIndex )
}

func ArrayDemo4()  {
	var chars [26] byte

	chars[0] = 'A'

	for i := 1; i < 26; i++ {
		chars[i]=chars[i-1]+1		
	}
	for index ,value := range chars {
		fmt.Printf("chars[%d]=%c \n",index,value)
	}
	
}

func ArrayDemo5()  {
	var arr [7] int
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(arr) - 1; i++ {
		radom := rand.Intn(100)
		arr[i] = radom
	}
	fmt.Println("随机数组arr=",arr)
   
	count := len(arr) 
    var arr2 [7] int
	for i := count - 1; i >= 0; i-- {

		fmt.Printf("反转打印数组arr[%d]=%d \n",i,arr[i])
        arr2[count -i-1]=arr[i]
	}
	fmt.Println("新数组arr2=",arr2)
	fmt.Println("原数组arr=",arr)
	for i := 0; i < count / 2; i++ {
		temp := arr[i]
		arr[i] = arr[ count - 1 - i ]
		arr[ count - 1 - i ] =temp
		fmt.Println("arr=",arr)
	}

}

/*冒泡排序*/

func BubbleSort()  {
	var arrys =[] int {24,69,80,57,30}
	flag :=false
	for i := 0; i < len(arrys) - 1; i++ {

		for j := 0; j < len(arrys)-1-i; j++ {
			if arrys[j] > arrys[j+ 1] {
				temp := arrys[j]
				arrys[j] = arrys[j+ 1]
				arrys[j + 1] = temp
				flag = true
			}
			fmt.Printf("第%d次外部排序中第%d次内部排序后array=%v \n",i+1,j+1,arrys)
		}
		fmt.Printf("第%d次外部排序后array=%v \n" ,i+1,arrys)
		if flag {
			flag = false
		}else {
			break
		}
	}
	fmt.Println("排序后的数组arry=",arrys)

}


func BinarySeach(arr *[6]int,startIndex int,lastIndex int,searchValue int)  {
	
	if startIndex > lastIndex {
		fmt.Println("对不起，找不到")
		return 
	}

	middleIndex :=( startIndex + lastIndex ) / 2

	if searchValue < arr[middleIndex] {
       BinarySeach(arr,startIndex,middleIndex - 1,searchValue)
	}else if searchValue > arr[middleIndex] {
		BinarySeach(arr,middleIndex + 1,lastIndex,searchValue)
	}else{
		fmt.Printf("二分查找的值%d的下标是%d \n",searchValue,middleIndex)
	}
}

// 定义二维数组，用于保存三个班，每个班五名同学成绩，
// 并求出每个班级平均分、以及所有班级平均分

func TwoArrays()  {
	var scores [3][5]float64
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("请输入第%d班的第%d学生的成绩\n", i + 1, j + 1)
			fmt.Scanln(&scores[i][j])
		}
	}

	//遍历二维数组求出总分
	totalSum := 0.0
	classNum :=  len(scores)
	for i := 0; i < classNum; i++ {

		//对每个班遍历
		sum := 0.0
		for j := 0; j < len(scores[0]); j++ {
			sum += scores[i][j]
		}
		totalSum += sum // 将所有班级的分数累计
		fmt.Printf("第%d班的平均分=%f\n", i+1, sum/float64(len(scores[i])))
	}
	   
		for i,v := range scores {
			sum := 0.0
			for _,value := range v {
				sum += value
			}
			fmt.Printf("第%d班的平均分=%f\n", i+1, sum/float64(len(scores[i])))
			
		}
	fmt.Printf("所有学生的平均分=%f\n", totalSum/float64(len(scores[0]) * classNum))

}
