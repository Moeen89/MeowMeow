package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	address = "localhost:5052"
)

// Create a random string with 20 length
func generateRandomText() string {
	// Set the desired length of the random text
	length := 20

	// Create a byte slice to hold the random bytes
	randomBytes := make([]byte, length)

	// Read random bytes into the slice
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println("Error generating random bytes:", err)
		return ""
	}

	// Convert the random bytes to a string
	randomText := fmt.Sprintf("%x", randomBytes)

	// Trim the string to the desired length
	randomText = randomText[:length]
	return randomText
}

// Generate a random even number between min and max (inclusive)
func generateRandomEven(min, max int64) (int64, error) {
	for {
		// Generate a random number between min and max
		randomNum, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
		if err != nil {
			return 0, err
		}

		// Add min to the random number to get the desired range
		randomNum.Add(randomNum, big.NewInt(min))

		// Make sure the number is even
		if randomNum.Int64()%2 == 0 {
			return randomNum.Int64(), nil
		}
	}
}

/*func main() {
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	replyMsg, err := c.ReqPq(ctx, &pb.Msg{
		Nonce:     clientNonce,
		MessageId: int32(messageId)})
	log.Printf("clientNonce: %s \n serverNonce: %s \n messageID: %d \n pNumber: %d \n gNumber: %d",
		replyMsg.GetNonce(), replyMsg.GetServerNonce(), replyMsg.GetMessageId(), replyMsg.GetP(), replyMsg.GetG())
}*/
