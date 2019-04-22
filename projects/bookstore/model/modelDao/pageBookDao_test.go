package modelDao

import (
	"fmt"
	"testing"
)

func TestPageBook(t *testing.T) {
	fmt.Println("测试PageBook中的函数")
	t.Run("测试分页Books：", testPageBook)

}



func testPageBook(t *testing.T) {
	
	page, _ := QueryPageBooksByBack(3,4)
	fmt.Println("获取书籍信息是：", page)
	for _, v := range page.Books {
		fmt.Println("图书的信息是：", v)
	}
	
}

