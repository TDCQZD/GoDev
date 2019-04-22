package model

import (
	"fmt"
	"goweb_code/web_DB/web_code/utils"
)

type User struct {
	Id   int
	Name string
	Age  string
	Sex  string
}

var db = utils.Db

//添加数据
func (user *User) AddUser() (err error) {

	//1.sql语句
	sqlStr := `insert into user (name, age, sex) values (?, ?, ?)`
	//2.执行
	res, err := db.Exec(sqlStr, user.Name, user.Age, user.Sex)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return
	}
	fmt.Println("执行结果：", res)
	return
}

//添加数据:Prepare预编译
func (user *User) AddUserWithPrepare() (err error) {
	//1.sql语句
	sqlStr := `insert into user (name, age, sex) values (?, ?, ?)`
	//2.预编译
	inStmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常：", err)
		return
	}
	//3.执行
	res, err := inStmt.Exec(user.Name, user.Age, user.Sex)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return
	}
	fmt.Println("执行结果：", res)
	return
}

//查询单条数据
func (user *User) SearchUser(id int) (*User, error) {
	sqlStr := `select id, name, age, sex from user where id = ?`
	// sqlStr := `select *from users where id = ?`

	row := db.QueryRow(sqlStr, id)

	err := row.Scan(&user.Id, &user.Name, &user.Age, &user.Sex)
	if err != nil {
		return nil, err
	}

	return user, nil
}

//查询所有数据
func (user *User) SearchAllUser() ([]*User, error) {
	sqlStr := `select id, name, age, sex from user `
	// sqlStr := `select *from users ?`

	//执行
	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	//创建User切片
	var users []*User
	for rows.Next() { //遍历数据
		userInfor := &User{}
		err := rows.Scan(&userInfor.Id, &userInfor.Name, &userInfor.Age, &userInfor.Sex)
		if err != nil {
			return nil, err
		}

		users = append(users, userInfor)
	}
	return users, nil
}

//修改数据
func (user *User) UpdateUser() (err error) {
	//1.sql语句
	sqlStr := `UPDATE user SET name = ?, age = ?, sex = ? where id = ?`
	//2.执行
	res, err := db.Exec(sqlStr, user.Name, user.Age, user.Sex, user.Id)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return
	}
	fmt.Println("执行结果：", res)
	return nil
}

//修改数据:Prepare预编译
func (user *User) UpdateUserWithPrepare() (err error) {
	//1.sql语句
	sqlStr := `UPDATE user SET name = ?, age = ?, sex = ? where id = ?`
	//2.预编译
	inStmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常：", err)
		return
	}
	//3.执行
	res, err := inStmt.Exec(user.Name, user.Age, user.Sex, user.Id)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return
	}
	fmt.Println("执行结果：", res)
	return nil
}

//删除数据
func (user *User) DeleteUser(id int) (err error) {
	//1.sql语句
	sqlStr := `DELETE FROM user WHERE id = ?`
	//2.执行
	res, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return
	}
	fmt.Println("执行结果：", res)
	return nil
}

//删除数据:Prepare预编译
func (user *User) DeleteUserWithPrepare(id int) (err error) {
	//1.sql语句
	sqlStr := `DELETE FROM user WHERE id = ?`
	//2.预编译
	inStmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常：", err)
		return
	}
	//3.执行
	res, err := inStmt.Exec(id)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return
	}
	fmt.Println("执行结果：", res)
	return nil
}
