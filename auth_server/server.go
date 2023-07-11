package main

import (
	"context"
	"crypto/rand"
	"crypto/sha1"
	pb "example.com/go-connection-grpc/GRPC"
	"fmt"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"log"
	"math/big"
	"net"
	"strconv"
	"time"
)

const (
	port = ":5052"
)

type connect struct {
	pb.UnimplementedAuthServiceServer
}

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

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

// Generate a random prime number between min and max (inclusive)
func generateRandomPrime(min, max int64) (int64, error) {
	for {
		// Generate a random number between min and max
		randomNum, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
		if err != nil {
			return 0, err
		}

		// Add min to the random number to get the desired range
		randomNum.Add(randomNum, big.NewInt(min))

		// Check if the generated number is prime
		if randomNum.ProbablyPrime(10) {
			return randomNum.Int64(), nil
		}
	}
}

// Generate a random odd number between min and max (inclusive)
func generateRandomOdd(min, max int64) (int64, error) {
	for {
		// Generate a random number between min and max
		randomNum, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
		if err != nil {
			return 0, err
		}

		// Add min to the random number to get the desired range
		randomNum.Add(randomNum, big.NewInt(min))

		// Make sure the number is odd
		if randomNum.Int64()%2 == 1 {
			return randomNum.Int64(), nil
		}
	}
}

func powInt(x, y, p int64) int32 {
	result := 1
	for i := 0; i < int(y); i++ {
		result = result * int(x)
		result = result % int(p)
	}
	return int32(result)
}

func (c *connect) ReqPq(ctx context.Context, in *pb.Msg) (*pb.ReplyMsg, error) {
	clientNonce := in.GetNonce()
	/// clientMessageId := in.GetMessageId()
	// TODO
	// serverNonce := generateRandomText()
	serverMessageId, err := generateRandomOdd(100, 100000)
	if err != nil {
		log.Printf("server messageId is 0!")
	}
	//pNumber, err := generateRandomPrime(100, 1000)
	//if err != nil {
	//	log.Printf("p_number is not set!")
	//}
	//gNumber, err := generateRandomPrime(3, 1000)
	//if err != nil {
	//	fmt.Println("g_number is not set", err)
	//}
	serverNonce := "Hossein"
	pNumber := 23
	gNumber := 5
	sha1String := clientNonce + serverNonce
	sha1Key := sha1.Sum([]byte(sha1String))
	sx16 := fmt.Sprintf("%x", sha1Key)
	session := map[string]string{"clientNonce": clientNonce, "serverNonce": serverNonce,
		"serverMessageId": strconv.FormatInt(serverMessageId, 10), "pNumber": strconv.FormatInt(int64(pNumber), 10),
		"gNumber": strconv.FormatInt(int64(gNumber), 10)}

	for k, v := range session {
		err := client.HSet(ctx, sx16, k, v).Err()
		if err != nil {
			panic(err)
		}
	}

	expiration := 20 * time.Minute // Example: Set expiration for 10 minutes
	err = client.Expire(context.TODO(), "myMap", expiration).Err()
	if err != nil {
		fmt.Println("Error setting expiration for map:", err)
	}

	return &pb.ReplyMsg{
		Nonce:       clientNonce,
		ServerNonce: serverNonce,
		MessageId:   int32(serverMessageId),
		P:           int32(pNumber),
		G:           int32(gNumber),
	}, nil
}

func (c *connect) Req_DHParam(ctx context.Context, in *pb.NewMsg) (*pb.NewReplyMsg, error) {
	clientNonce := in.GetNonce()
	//clientMessageId := in.GetMessageId()
	serverNonce := in.GetServerNonce()
	sha1String := clientNonce + serverNonce
	sha1Key := sha1.Sum([]byte(sha1String))
	sx16 := fmt.Sprintf("%x", sha1Key)
	userSession := client.HGetAll(ctx, sx16).Val()
	// TODO
	//randomNumber, err := generateRandomOdd(10, 200)
	serverMessageId, err := generateRandomOdd(100, 100000)
	if err != nil {
		log.Printf("server messageId is 0!")
	}
	g, err := strconv.ParseInt(userSession["gNumber"], 10, 32)
	p, err := strconv.ParseInt(userSession["pNumber"], 10, 32)
	a := in.GetANumber()
	b := powInt(g, 15, p) % int32(p)
	publicKey := powInt(int64(a), 15, p) % int32(p)
	client.Set(ctx, strconv.Itoa(int(publicKey)), "ok", 0)
	client.Del(ctx, sx16)

	if err != nil {
		panic(err)
	}
	return &pb.NewReplyMsg{
		Nonce:       clientNonce,
		ServerNonce: serverNonce,
		MessageId:   int32(serverMessageId),
		BNumber:     b,
	}, nil
}

func (c *connect) CheckKey(ctx context.Context, in *pb.Key) (*pb.Val, error) {
	key := in.GetAuthKey()
	return &pb.Val{
		IsTrue: int32(client.Exists(ctx, strconv.Itoa(int(key))).Val()),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &connect{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
