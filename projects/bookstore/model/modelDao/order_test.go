package modelDao

import (
	"fmt"

	"testing"
)

func TestOrders(t *testing.T) {
	fmt.Println("测试bookdao中的函数")

	t.Run("测试修改Books：", testQueryOrderItems)

}

func testQueryOrderItems(t *testing.T) {
	orderItems, _ := SearchOrderItems("49353361-3a1c-4753-4478-872f2e32f995")
	for _,v := range orderItems {
		fmt.Println("获取订单信息是：", v.Book.ID)
	}
	
}

