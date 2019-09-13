package main

import "fmt"

/*
type slice struct {
	array unsafe.Pointer  // 内部是分配一个连续的内存块，这个指针就指向这个内存的首地址
	len   int // 数组长度
	cap   int // 数组容量
}
当不断的向slice添加数据，len就不断变大，当len > cap的时候，内部开始扩容，重新分配一块内存，大致的扩容策略是cap * 2，然后将老地址的数据copy到新地址。

由于以上的特性，当有一个slice a, 将a 复制给b, 这个时候，a, b指向同一个底层数组，当对a不断的进行添加操作，当a进行扩容后，a,b便开始指向不同的底层数组。所有在操作slice的时候要格外注意。

*/

func makeSlice() {
	s1 := make([]int, 0, 2)
	s2 := make([]int, 2, 2)
	s3 := []int{}
	s4 := []int{1, 2}
	fmt.Printf("s1: %v %p;    len:%v;    cap:%v     \n", s1, s1, len(s1), cap(s1))
	fmt.Printf("s2: %v %p;    len:%v;    cap:%v     \n", s2, s2, len(s2), cap(s2))
	fmt.Printf("s3: %v %p;    len:%v;    cap:%v     \n", s3, s3, len(s3), cap(s3))
	fmt.Printf("s4: %v %p;    len:%v;    cap:%v     \n", s4, s4, len(s4), cap(s4))
	fmt.Println("----------------------------------------------------------")
	s1 = append(s1, 3, 4)
	s2 = append(s2, 3, 4)
	s3 = append(s3, 3, 4)
	s4 = append(s4, 3, 4)
	fmt.Printf("s1: %v %p;    len:%v;    cap:%v     \n", s1, s1, len(s1), cap(s1))
	fmt.Printf("s2: %v %p;    len:%v;    cap:%v     \n", s2, s2, len(s2), cap(s2))
	fmt.Printf("s3: %v %p;    len:%v;    cap:%v     \n", s3, s3, len(s3), cap(s3))
	fmt.Printf("s4: %v %p;    len:%v;    cap:%v     \n", s4, s4, len(s4), cap(s4))
	fmt.Println("----------------------------------------------------------")
	for _, v := range s4 {
		fmt.Printf("range: ")
		fmt.Println(v)
	}
	for i := 0; i < len(s4); i++ {
		fmt.Printf("for: ")
		fmt.Println(s4[i])
	}
}

type V struct {
	i int
}
type W struct {
	v []*V
}

func strcutSlice() {
	s := W{}.v
	fmt.Printf("s: %v %p;    len:%v;    cap:%v     \n", s, s, len(s), cap(s))
	v1 := &V{1}
	v2 := &V{2}
	s = append(s, v1, v2)
	fmt.Printf("s: %v %p;    len:%v;    cap:%v     \n", s, s, len(s), cap(s))
	for _, v := range s {
		fmt.Printf("range: ")
		fmt.Println(v.i)
	}
}
func main() {
	// makeSlice()
	strcutSlice()
}
