package linkpool

import (
	"fmt"
	"log"
	"github.com/garyburd/redigo/redis"
)


//错误检查
func CheckError(err error) {
    if err != nil {
		log.Fatal(err)
        panic(err)
    }
}

var pool *redis.Pool

func init()  {	
	pool = &redis.Pool{
				MaxIdle : 8,
				MaxActive : 0,
				IdleTimeout : 10,
				Dial : func() (redis.Conn, error) {return redis.Dial("tcp","localhost:6379")},
	}
}

func main()  {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("set","pool","链接池")
	CheckError(err)

	res, err := redis.String(conn.Do("get","pool"))
	CheckError(err)
	fmt.Println(res)
	pool.Close()
}