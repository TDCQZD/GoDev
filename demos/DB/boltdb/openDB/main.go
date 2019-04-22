package main

import (
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	//打开当前目录中的my.db数据文件。 如果它不存在，它将被创建。
	_, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	
}
