package modelDao

import (
	"fmt"
	"goweb_code/bookstore/DB"
	"goweb_code/bookstore/model"
)

var sessionDB = DB.DB

//添加Session数据
func AddSession(session *model.Seesion) error {

	//1.sql语句
	sqlStr := `INSERT INTO session (UUID, USERID, USERNAME) values (?,?,?)`
	//2.执行
	res, err := sessionDB.Exec(sqlStr, session.UUID, session.UserID, session.UserName)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	fmt.Println("执行结果：", res)
	return nil
}

//查询session
func SearchSessionByUUID(uuid string) (*model.Seesion, error) {
	sqlStr := `SELECT  UUID, USERID, USERNAME FROM session WHERE UUID = ?`

	row := sessionDB.QueryRow(sqlStr, uuid)
	session := &model.Seesion{}
	err := row.Scan(&session.UUID, &session.UserID, &session.UserName)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// 删除Session数据
func DeleteSession(uuid string) error {
	//1.sql语句
	sqlStr := `DELETE FROM session WHERE UUID = ?`
	//2.执行
	res, err := sessionDB.Exec(sqlStr, uuid)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	fmt.Println("执行结果：", res)
	return nil
}
