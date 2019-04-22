package modelDao

import (
	"fmt"
	"goweb_code/bookstore/model"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("测试userdao中的函数")
	// t.Run("测试添加：", testUserAdd)
	// t.Run("测试修改：", testUserUpdate)
	// t.Run("测试查询：", testUserQuery)
	// t.Run("测试删除：", testUserDelete)
	// t.Run("测试登录：", testLogin)
	t.Run("测试注册：", testRegister)

}
func testUserAdd(t *testing.T) {
	user := &model.Users{
		Username: "王五",
		Password: "111111",
		Email:    "ww@qq.com",
	}
	err := UserAdd(user)
	fmt.Println("结果是：", err)
}

func testUserUpdate(t *testing.T) {
	user := &model.Users{
		Username: "王五",
		Password: "123456",
		Email:    "ww@126.com",
	}
	err := UserUPdate(user)
	fmt.Println("结果是：", err)
}

func testUserQuery(t *testing.T) {
	user, err := UserQueryByUserName("王五")
	fmt.Println("结果是：", err, user.Username)
}

func testUserDelete(t *testing.T) {
	err := UserDelete("王五")
	fmt.Println("结果是：", err)
}

func testLogin(t *testing.T) {
	flag, _, err := LoginAccount("王1五", "1223456")
	fmt.Println("获取用户登录信息是：", flag, err)
}

func testRegister(t *testing.T) {
	user := &model.Users{
		Username: "赵四",
		Password: "123456",
		Email:    "1243@11.com",
	}
	flag, err := RegisterAccount(user)
	fmt.Println("获取用户注册信息是：", flag, err)
}
