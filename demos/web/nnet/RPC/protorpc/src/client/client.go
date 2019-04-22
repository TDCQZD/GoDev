package main

import (
	"fmt"
	pb "goCode/nnet/RPC/protorpc/src/pb"
	"log"
)

func main() {
	conn, err := pb.DialArithService("tcp", "127.0.0.1:8097")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}
	defer conn.Close()

	req := &pb.ArithRequest{9, 2}

	res, err := conn.Multiply(req)
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d * %d = %d\n", req.GetA(), req.GetB(), res.GetPro())

	res, err = conn.Divide(req)
	if err != nil {
		log.Fatalln("arith error ", err)
	}
	fmt.Printf("%d / %d, quo is %d, rem is %d\n", req.A, req.B, res.Quo, res.Rem)
}
