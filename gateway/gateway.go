package main

import (
	"awesomeProject2/docs"
	"context"
	"fmt"
	"github.com/JGLTechnologies/gin-rate-limit"
	"github.com/redis/go-redis/v9"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
	"google.golang.org/grpc"

	pb "awesomeProject2/GRPC"
	pbs "awesomeProject2/pb"

	"github.com/gin-gonic/gin"
)

const (
	address    = "localhost:5052"
	addressBiz = "localhost:5062"
)

type req_req_pq struct {
	Id    string `json:"messageId" form:"messageId"`
	Nonce string `json:"nonce" form:"nonce"`
}

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	//c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
	client.Set(c, "ip"+c.ClientIP(), "banned", 24*time.Hour)
}

func startServer() {
	router := gin.Default()
	// This makes it so each ip can only make 100 requests per second
	store := ratelimit.RedisStore(&ratelimit.RedisOptions{
		RedisClient: client,
		Rate:        time.Second,
		Limit:       100,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})
	// Define the routes for the API Gateway
	router.Any("/AUTH/*path", creatLim(), mw, auth)
	router.Any("/BIZ/*path", biz)
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", auth)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// Start the API Gateway
	router.Run(":6433")
}
func creatLim() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		if client.Exists(c, "ip"+ip).Val() == 1 {
			c.String(429, "Too many requests. banned for 24h")
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		} else {
			c.Next()
		}

	}
}

// @BasePath /AUTH

// @Summary handle authentication
// @Schemes
// @Description handle req_pq and req_DHparam of auth serveer, by receive http as input, and connecting to auth server and using grpc and return result
// @Tags Auth
// @Accept json
// @Produce json
// @Success 202 {string} servernonce and p,q
// @Router /gateway/auth [get]
func auth(c *gin.Context) {
	// Parse the target URL
	fmt.Print(c.Param("path"))
	// Create the reverse proxy
	//proxy := httputil.NewSingleHostReverseProxy(targetURL)
	if c.Param("path") == "/req_pq" {
		var r req_req_pq
		if err := c.Bind(&r); err != nil {
			fmt.Println(err)
		}
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer func(conn *grpc.ClientConn) {
			err := conn.Close()
			if err != nil {

			}
		}(conn)
		c2 := pb.NewAuthServiceClient(conn)
		messageId, _ := strconv.Atoi(r.Id)
		clientNonce := r.Nonce
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		replyMsg, err := c2.ReqPq(ctx, &pb.Msg{
			Nonce:     clientNonce,
			MessageId: int32(messageId)})
		if replyMsg == nil {
			fmt.Print("can't connect\n")
		}

		c.JSON(http.StatusAccepted, gin.H{
			"nonce": replyMsg.GetNonce(), "server_nonce": replyMsg.GetServerNonce(), "message_id": replyMsg.GetMessageId(), "p": replyMsg.GetP(), "g": replyMsg.GetG(),
		})
		return
	} else if c.Param("path") == "/req_DH_params" {
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer func(conn *grpc.ClientConn) {
			err := conn.Close()
			if err != nil {

			}
		}(conn)
		c2 := pb.NewAuthServiceClient(conn)
		jsonData, _ := ioutil.ReadAll(c.Request.Body)
		messageIdt := gjson.Get(string(jsonData), "messageId")
		messageId, _ := strconv.Atoi(messageIdt.Str)
		clientNonce := gjson.Get(string(jsonData), "clientNonce").Str
		serverNonce := gjson.Get(string(jsonData), "serverNonce").Str
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		//randomNumber, err := generateRandomEven(10, 200)
		at := gjson.Get(string(jsonData), "a")
		a, _ := strconv.Atoi(at.Str)

		defer cancel()
		newReplyMsg, err := c2.Req_DHParam(ctx, &pb.NewMsg{
			Nonce:       clientNonce,
			ServerNonce: serverNonce,
			MessageId:   int32(messageId),
			ANumber:     int32(a)})
		c.JSON(http.StatusAccepted, gin.H{
			"b": newReplyMsg.GetBNumber(),
		})
		return
	}

}

func biz(c *gin.Context) {
	// Parse the target URL

	// Create the reverse proxy
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	//proxy := httputil.NewSingleHostReverseProxy(targetURL)
	authToken := gjson.Get(string(jsonData), "authKey")
	if authToken.Str == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "AUTH require",
		})
		return
		//send auth token to auth server and check token validity
	}
	if c.Param("path") == "/get_users" {
		conn, err := grpc.Dial(addressBiz, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c2 := pbs.NewGetUsersClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
		defer cancel()
		userId := gjson.Get(string(jsonData), "userId")
		idx, _ := strconv.Atoi(userId.Str)
		idd := int32(idx)
		messageId := gjson.Get(string(jsonData), "messageId")
		messageid, _ := strconv.Atoi(messageId.Str)
		message_id := int32(messageid)
		authKey := authToken.Str
		authkey, _ := strconv.Atoi(authKey)
		authkey32 := int32(authkey)

		req := &pbs.UserRequest{UserId: idd, MessageId: message_id, AuthKey: authkey32}
		r, err := c2.GetUsers(ctx, req)
		if err != nil {
			log.Printf("could not get users:  %v", err)
			c.JSON(http.StatusAccepted, gin.H{
				"error": "could not get users",
			})
			return
		} else {
			var x []pbs.User
			for _, user := range r.GetUsers() {
				x = append(x, *user)
			}

			c.JSON(http.StatusAccepted, gin.H{
				"messageId": r.GetMessageId(), "users": x,
			})
			return
		}
	}
	if c.Param("path") == "/get_users_with_sql_inject" {
		conn, err := grpc.Dial(addressBiz, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c2 := pbs.NewGetUsersClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
		defer cancel()
		userId := gjson.Get(string(jsonData), "userId")
		messageId := gjson.Get(string(jsonData), "messageId")
		messageid, _ := strconv.Atoi(messageId.Str)
		message_id := int32(messageid)
		authKey := authToken.Str
		authkey, _ := strconv.Atoi(authKey)
		authkey32 := int32(authkey)

		req := &pbs.UserRequestWithSqlInject{UserId: userId.Str, MessageId: message_id, AuthKey: authkey32}
		r, err := c2.GetUsersWithSqlInject(ctx, req)
		if err != nil {
			log.Printf("could not get users:  %v", err)
			c.JSON(http.StatusAccepted, gin.H{
				"error": "auth failed",
			})
			return
		} else {
			var x []pbs.User
			for _, user := range r.GetUsers() {
				x = append(x, *user)
			}

			c.JSON(http.StatusAccepted, gin.H{
				"messageId": r.GetMessageId(), "users": x,
			})
			return
		}
	}
}

func main() {
	startServer()
}
