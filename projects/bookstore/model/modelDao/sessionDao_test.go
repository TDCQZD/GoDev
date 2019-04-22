package modelDao

import (
	"fmt"
	"goweb_code/bookstore/model"
	"goweb_code/bookstore/utils"
	"testing"
)

func TestSession(t *testing.T) {
	fmt.Println("测试Session中的函数")
	// t.Run("测试Session 添加：", testAddSession)
	t.Run("测试Session 删除：", testDeleteSession)
	// t.Run("测试Session 查询：", testQuerySession)
}

func testAddSession(t *testing.T) {
	uuid := utils.CreateUUID()
	session := model.Seesion{
		UUID:     uuid,
		UserID:   "7c6b6ded-0dc1-4171-660c-9cbb21e0f206",
		UserName: "qqqq",
	}
	err := AddSession(&session)
	fmt.Println("添加session：", err)
}
func testQuerySession(t *testing.T) {

	seeion, _ := SearchSessionByUUID("8eefffd2-a861-4278-559a-009808be71d5")
	fmt.Println("查询session：", seeion)
}

func testDeleteSession(t *testing.T) {

	err := DeleteSession("1")
	fmt.Println("删除ssession：", err)
}
