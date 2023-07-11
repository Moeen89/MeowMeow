package main

import (
	"awesomeProject2/docs"
	"context"
	"fmt"
	"github.com/JGLTechnologies/gin-rate-limit"
	"github.com/redis/go-redis/v9"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"strconv"
	"time"

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
type req_req_DH_params struct {
	Id          string `json:"messageId" form:"messageId"`
	ClientNonce string `json:"clientNonce" form:"clientNonce"`
	ServerNonce string `json:"serverNonce" form:"serverNonce"`
	A           string `json:"a" form:"a"`
}

type req_get_users struct {
	Id      string `json:"messageId" form:"messageId"`
	AuthKey string `json:"authKey" form:"authKey"`
	UserId  string `json:"userId" form:"userId"`
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
	router.Any("/AUTH/req_pq", creatLim(), mw, req_pq)
	router.Any("/AUTH/req_DH_params", creatLim(), mw, req_DH_params)
	router.Any("/BIZ/get_users", get_users)
	router.Any("/BIZ/get_users_with_sql_inject", get_users_with_sql_inject)
	docs.SwaggerInfo.BasePath = ""
	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", req_pq)
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
// @Router /AUTH/req_pq [post]
// @param	data      body     req_req_pq     false  "nonce and message id for pq"
func req_pq(c *gin.Context) {
	// Parse the target URL
	// Create the reverse proxy
	//proxy := httputil.NewSingleHostReverseProxy(targetURL)

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

}

// @BasePath /AUTH
// @Summary handle authentication
// @Schemes
// @Description handle req_pq and req_DHparam of auth serveer, by receive http as input, and connecting to auth server and using grpc and return result
// @Tags Auth
// @Accept json
// @Produce json
// @Success 202 {string} servernonce and p,q
// @Router /AUTH/req_DH_params [post]
// @param	data      body     req_req_DH_params     false  "nonce and message id for pq"
func req_DH_params(c *gin.Context) {
	var r req_req_DH_params
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//randomNumber, err := generateRandomEven(10, 200)
	a, _ := strconv.Atoi(r.A)

	defer cancel()
	newReplyMsg, err := c2.Req_DHParam(ctx, &pb.NewMsg{
		Nonce:       r.ClientNonce,
		ServerNonce: r.ServerNonce,
		MessageId:   int32(messageId),
		ANumber:     int32(a)})
	c.JSON(http.StatusAccepted, gin.H{
		"b": newReplyMsg.GetBNumber(),
	})
}

// @BasePath /AUTH
// @Summary handle authentication
// @Schemes
// @Description handle req_pq and req_DHparam of auth serveer, by receive http as input, and connecting to auth server and using grpc and return result
// @Tags Auth
// @Accept json
// @Produce json
// @Success 202 {string} servernonce and p,q
// @Router /BIZ/get_users [post]
// @param	data      body     req_get_users     false  "nonce and message id for pq"
func get_users(c *gin.Context) {
	var r req_get_users
	if err := c.Bind(&r); err != nil {
		fmt.Println(err)
	}
	authToken := r.AuthKey
	if authToken == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "AUTH require",
		})
		return
		//send auth token to auth server and check token validity
	}
	conn, err := grpc.Dial(addressBiz, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c2 := pbs.NewGetUsersClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	idx, _ := strconv.Atoi(r.UserId)
	idd := int32(idx)
	messageid, _ := strconv.Atoi(r.Id)
	message_id := int32(messageid)
	authKey := authToken
	authkey, _ := strconv.Atoi(authKey)
	authkey32 := int32(authkey)

	req := &pbs.UserRequest{UserId: idd, MessageId: message_id, AuthKey: authkey32}
	rc, err := c2.GetUsers(ctx, req)
	if err != nil {
		log.Printf("could not get users:  %v", err)
		c.JSON(http.StatusAccepted, gin.H{
			"error": "could not get users",
		})
		return
	} else {
		var x []pbs.User
		for _, user := range rc.GetUsers() {
			x = append(x, *user)
		}

		c.JSON(http.StatusAccepted, gin.H{
			"messageId": rc.GetMessageId(), "users": x,
		})
		return
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
// @Router /BIZ/get_users_with_sql_inject [post]
// @param	data      body     req_get_users     false  "nonce and message id for pq"
func get_users_with_sql_inject(c *gin.Context) {
	var r req_get_users
	if err := c.Bind(&r); err != nil {
		fmt.Println(err)
	}
	authToken := r.AuthKey
	if authToken == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "AUTH require",
		})
		return
		//send auth token to auth server and check token validity
	}
	conn, err := grpc.Dial(addressBiz, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c2 := pbs.NewGetUsersClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	userId := r.UserId
	messageid, _ := strconv.Atoi(r.Id)
	message_id := int32(messageid)
	authKey := authToken
	authkey, _ := strconv.Atoi(authKey)
	authkey32 := int32(authkey)

	req := &pbs.UserRequestWithSqlInject{UserId: userId, MessageId: message_id, AuthKey: authkey32}
	rc, err := c2.GetUsersWithSqlInject(ctx, req)
	if err != nil {
		log.Printf("could not get users:  %v", err)
		c.JSON(http.StatusAccepted, gin.H{
			"error": "auth failed",
		})
		return
	} else {
		var x []pbs.User
		for _, user := range rc.GetUsers() {
			x = append(x, *user)
		}

		c.JSON(http.StatusAccepted, gin.H{
			"messageId": rc.GetMessageId(), "users": x,
		})
		return
	}

}

func main() {
	startServer()
}
