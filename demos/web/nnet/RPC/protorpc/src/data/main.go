package main

import (
	pb "goCode/nnet/RPC/protorpc/src/pb"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	user1 := pb.User{
		Id:   *proto.Int32(1),
		Name: *proto.String("Mike"),
	}

	user2 := pb.User{
		Id:   2,
		Name: "John",
	}

	users := pb.MultiUser{
		Users: []*pb.User{&user1, &user2},
	}

	// 序列化数据
	data, err := proto.Marshal(&users)
	if err != nil {
		log.Fatalln("Marshal data error: ", err)
	}
	println(users.Users[0].GetName()) // output: Mike

	// 对已序列化的数据进行反序列化
	var target pb.MultiUser
	err = proto.Unmarshal(data, &target)
	if err != nil {
		log.Fatalln("Unmarshal data error: ", err)
	}
	println(target.GetUsers()[1].Name) // output: John

}
