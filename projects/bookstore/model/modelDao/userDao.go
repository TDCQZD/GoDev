package modelDao

import (
	"errors"
	"goweb_code/bookstore/DB"
	"goweb_code/bookstore/model"
)

var userDB = DB.DB

//验证用户名
func VerifyName(username string) (bool, error) {
	_, err := UserQueryByUserName(username)
	if err != nil {
		err = errors.New("该用户不存在！")
		return false, err
	}
	return true, nil
}


//用户注册 添加
func RegisterAccount(user *model.Users) (bool, error) {
	_, err := UserQueryByUserName(user.Username)
	if err == nil { //用户已存在，注册失败
		err = errors.New("该用户已存在！")
		return false, err
	}
	err = UserAdd(user)
	if err != nil {
		return false, err
	}

	return true, nil
}

//用户登录 查询
func LoginAccount(username string, password string) (bool, *model.Users, error) {
	flag, user, err := verifyAccount(username, password)
	// flag, err = verifyNameAndPwd1(username, password)
	return flag, user, err
}

func UserAdd(user *model.Users) error {
	//1.sql语句
	sqlStr := `INSERT INTO shopusers (ID,NAME,PWD,EMAIL) VALUES (?,?,?,?)`
	//2.执行
	_, err := userDB.Exec(sqlStr, user.GetUserID(), user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func UserUPdate(user *model.Users) error {

	//1.sql语句
	sqlStr := `UPDATE  shopusers SET PWD = ? ,EMAIL = ?  WHERE NAME = ?`
	//2.执行
	_, err := userDB.Exec(sqlStr, user.Password, user.Email, user.Username)
	if err != nil {
		return err
	}
	return nil
}

func UserDelete(username string) error {
	//1.sql语句
	sqlStr := `DELETE FROM shopusers WHERE NAME = ? `
	//2.执行
	_, err := userDB.Exec(sqlStr, username)
	if err != nil {
		return err
	}
	return nil
}

func UserQueryByUserName(username string) (*model.Users, error) {

	//1.sql语句
	sqlStr := `SELECT  ID, NAME, PWD, EMAIL FROM shopusers WHERE NAME = ? `
	row := userDB.QueryRow(sqlStr, username)
	user := &model.Users{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

/*登录前验证账户信息*/
func verifyAccount(username string, password string) (bool, *model.Users, error) {
	user, err := UserQueryByUserName(username)
	if err != nil { //用户不存在
		err = errors.New("该用户不存在！")
		return false, nil, err
	}

	if user.Password != password {
		err = errors.New("密码错误")
		return false, nil, err
	}
	return true, user, nil
}
