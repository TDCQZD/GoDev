package main

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"fmt"
	
)

func Client()  {
	//1. 链接到redis
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println("连接成功！")
	fmt.Println(conn)
	
	//2、通过go 向redis写入数据 string类型数据
	res, err := conn.Do("set","nickname","乱世狂刀")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("res:",res)

	//3. 通过go 向redis读取数据string类型数据
	name, err := redis.String(conn.Do("get","nickname"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name:",name)

	// 4. 通过go 向redis写入数据 Hash类型数据
	res1, err := conn.Do("hmset","user","name","一页书","age","100","skill","八部天龙火")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("res1:",res1)

	//5. 通过go 向redis读取数据 Hash类型数据
	user, err := redis.StringMap(conn.Do("hgetall","user"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("user:",user)
}

func main()  {
	Client()
}