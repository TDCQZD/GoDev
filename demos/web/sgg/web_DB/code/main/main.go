package main

import (
	"go_code/grammar/DB/mysql/code/bean"
	
	"io/ioutil"
	"strings"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
)

//数据库配置常量
const (
    userName = "root"
    password = "root"
    ip = "127.0.0.1"
    port = "3306"
	dbName = "godb"
	tableName = "./table.sql"
)
//链接数据库
func openDB() *sql.DB {
	
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
    //打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	db, err := sql.Open("mysql", path)
	
    //DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
    // db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/godb?charset=utf8");
   //验证连接err := DB.Ping();
	if  err != nil {
		fmt.Println("open database fail")
		log.Fatal(err)
	
    }
    //关闭数据库，db会被多个goroutine共享，可以不调用
	// defer db.Close();
	
	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)
	
	fmt.Println("connnect success")
	// fmt.Println(db)
	return db
}

//创建数据库表
func createTable(DB *sql.DB)  error {

    sqlBytes, err := ioutil.ReadFile(tableName);
    if err != nil {
        return err
    }
    sqlTable := string(sqlBytes);
    fmt.Println(sqlTable)
    
    _, err = DB.Exec(sqlTable)
    if err != nil {
        return err
    }
    return nil
}

//插入表数据
func InsertData(DB *sql.DB,user bean.User) bool {
    //开启事务
    tx, err := DB.Begin()
    if err != nil{
        fmt.Println("tx fail")
        return false
    }
    //准备sql语句
    stmt, err := tx.Prepare("INSERT INTO user (`id`, `name`,`age`,`sex`,`likes`) VALUES (?, ?, ?, ?, ?)")
    if err != nil{
        fmt.Println("Prepare fail")
        return false
    }
    //将参数传递到sql语句中并且执行
    res, err := stmt.Exec(user.Id, user.Name, user.Age,user.Sex,user.Likes)
    if err != nil{
        fmt.Println("Exec fail")
        return false
    }
    //将事务提交
    tx.Commit()
    //获得上一个插入自增的id
    fmt.Println(res.LastInsertId())
    return true
}

//查询所有数据
func SelectAllData(DB *sql.DB) ([]bean.User) {
    //执行查询语句
    rows, err := DB.Query("SELECT * from user")
    if err != nil{
        fmt.Println("SELECT fail")    
    }
    var users []bean.User
    //循环读取结果
    for rows.Next(){
        var user bean.User
        //将每一行的结果都赋值到一个user对象中
        err := rows.Scan(&user.Id, &user.Name, &user.Age,&user.Sex,&user.Likes)
        if err != nil {
            fmt.Println("rows fail")
        }
        //将user追加到users的这个数组中
        users = append(users, user)
    }
    return users
}

//查询单条数据
func SelectDataById(DB *sql.DB, id int) (bean.User) {
    var user bean.User
    err := DB.QueryRow("SELECT * FROM user WHERE id = ?", id).Scan(&user.Id, &user.Name, &user.Age,&user.Sex,&user.Likes)
    if err != nil{
        fmt.Println("查询出错了")
    }
    return user
}


// 修改数据
func UpdateData(DB *sql.DB, user bean.User) (bool) {
    //开启事务
    tx, err := DB.Begin()
    if err != nil{
        fmt.Println("tx fail")
    }
    //准备sql语句
    stmt, err := tx.Prepare("UPDATE user SET name = ?, age = ?, sex = ?, likes = ? WHERE id = ?")
    if err != nil{
        fmt.Println("Prepare fail")
        return false
    }
    //设置参数以及执行sql语句
    res, err := stmt.Exec(user.Name, user.Age, user.Sex, user.Likes, user.Id)
    if err != nil{
        fmt.Println("Exec fail")
        return false
    }
    //提交事务
    tx.Commit()
    fmt.Println(res.LastInsertId())
    return true
}


//删除数据
func DeleteData(DB *sql.DB, user bean.User) (bool) {
    //开启事务
    tx, err := DB.Begin()
    if err != nil{
        fmt.Println("tx fail")
    }
    //准备sql语句
    stmt, err := tx.Prepare("DELETE FROM user WHERE id = ?")
    if err != nil{
        fmt.Println("Prepare fail")
        return false
    }
    //设置参数以及执行sql语句
    res, err := stmt.Exec(user.Id)
    if err != nil{
        fmt.Println("Exec fail")
        return false
    }
    //提交事务
    tx.Commit()
    //获得上一个insert的id
    fmt.Println(res.LastInsertId())
    return true
}



func main()  {
	db := openDB()
	fmt.Println(db)
	// err := createTable(db)
	// if err != nil {//数据库表创建失败
    //    log.Fatal(err)
	// }
	// user := bean.User{1,"张三",20,"男","足球",}
	// InsertData(db,user)
	// res1 := SelectAllData(db)
	// fmt.Println(res1)
	// res2 := SelectDataById(db,3)
	// fmt.Println(res2)
	// user2 := bean.User{2,"王红",30,"女","购物",}
	// UpdateData(db,user2)
	// res2 = SelectDataById(db,2)
	// fmt.Println(res2)
	user3 := bean.User{Id:3,}
	DeleteData(db,user3)

}

