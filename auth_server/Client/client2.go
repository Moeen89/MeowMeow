package main

import (
	"context"
	pb "example.com/go-connection-grpc/GRPC"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func powInt(x, y, p int64) int32 {
	result := 1
	for i := 0; i < int(y); i++ {
		result = result * int(x)
		result = result % int(p)
	}
	return int32(result)
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	c := pb.NewAuthServiceClient(conn)
	messageId, err := generateRandomEven(100, 100000)
	// TODO
	// clientNonce := generateRandomText()
	clientNonce := "Amir"
	serverNonce := "Hossein"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//randomNumber, err := generateRandomEven(10, 200)
	a := powInt(5, 6, 23) % int32(23)

	defer cancel()
	newReplyMsg, err := c.Req_DHParam(ctx, &pb.NewMsg{
		Nonce:       clientNonce,
		ServerNonce: serverNonce,
		MessageId:   int32(messageId),
		ANumber:     a})
	b := newReplyMsg.GetBNumber()
	publicKey := powInt(int64(b), 6, 23) % int32(23)
	fmt.Println(publicKey)
}
