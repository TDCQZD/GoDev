package model

import (
	"fmt"
	"testing"
)

/**/
//TestMain函数可以在测试函数执行之前做一些其他操作
func TestMain(m *testing.M) {
	fmt.Println("测试开始：")
	//通过m.Run()来执行测试函数
	m.Run()
}

func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的相关方法")
	//通过t.Run()来执行子测试函数
	// t.Run("测试添加用户:", testAddUser)
	// t.Run("测试单条查询:", testGetUserByID)
	t.Run("测试多条查询:", testGetUsers)
	// t.Run("测试修改用户:", testUpdatUser)
	// t.Run("测试删除用户:", testDeleteUser)

}

//如果函数名不是以Test开头，那么该函数默认不执行，我们可以将它设置成为一个子测试函数
func testAddUser(t *testing.T) {
	fmt.Println("子测试函数执行：")
	user := &User{
		Name: "韩雪",
		Age:  "30",
		Sex:  "女",
	}
	//调用添加用户的方法
	// user.AddUser()
	user.AddUserWithPrepare()
}

func testUpdatUser(t *testing.T) {
	fmt.Println("子测试函数执行：")
	user := &User{
		Id:   3,
		Name: "王五",
		Age:  "20",
		Sex:  "男",
	}
	//调用修改用户的方法
	// user.UpdateUser()
	user.UpdateUserWithPrepare()
}

func testDeleteUser(t *testing.T) {
	fmt.Println("子测试函数执行：")
	user := &User{}
	//调用添加用户的方法
	// user.DeleteUser(5)
	user.DeleteUserWithPrepare(6)
}

//测试获取一个User
func testGetUserByID(t *testing.T) {
	fmt.Println("测试查询一条记录：")
	user := &User{}
	//调用获取User的方法
	u, _ := user.SearchUser(1)
	fmt.Println("得到的User的信息是：", u)
}

//测试获取所有User
func testGetUsers(t *testing.T) {
	fmt.Println("测试查询所有记录：")
	user := &User{}
	//调用获取所有User的方法
	us, _ := user.SearchAllUser()
	//遍历切片
	for k, v := range us {
		fmt.Printf("第%v个用户是:%v\n", k+1, v)
	}

}
