package main

import (
	"awesomeProject2/ratelimit"
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}

func createRouter() *gin.Engine {
	router := gin.Default()
	// This makes it so each ip can only make 5 requests per second
	store := ratelimit.RedisStore(&ratelimit.RedisOptions{
		RedisClient: redis.NewClient(&redis.Options{
			Addr: "localhost:7680",
		}),
		Rate:  time.Second,
		Limit: 100,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})
	// Define the routes for the API Gateway
	router.Any("/AUTH/*path", mw, createReverseProxyAuth("http://localhost:8081"))
	router.Any("/BIZ/*path", createReverseProxyBiz("http://localhost:8082"))
	return router
}

func startServer(router *gin.Engine) {
	// Start the API Gateway
	router.Run(":8080")
}
func createReverseProxyAuth(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the target URL
		targetURL, _ := url.Parse(target)

		// Create the reverse proxy
		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		if c.Param("path") == "req_pq" {
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
			messageIdt, _ := c.Get("messageId")
			messageId, _ := strconv.Atoi(messageIdt.(string))
			clientNonce, _ := c.Get("nonce")
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			replyMsg, err := c2.ReqPq(ctx, &pb.Msg{
				Nonce:     clientNonce.(string),
				MessageId: int32(messageId)})
			c.JSON(http.StatusAccepted, gin.H{
				"nonce": replyMsg.GetNonce(), "server_nonce": replyMsg.GetServerNonce(), "message_id": replyMsg.GetMessageId(), "p": replyMsg.GetP(), "g": replyMsg.G,
			})
			return
		} else if c.Param("path") == "req_DH_params" {
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
			messageIdt, _ := c.Get("messageId")
			messageId, _ := strconv.Atoi(messageIdt.(string))
			clientNonce := "Amir"
			serverNonce := "Hossein"
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			//randomNumber, err := generateRandomEven(10, 200)
			at, _ := c.Get("a")
			a, _ := strconv.Atoi(at.(string))

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

		// Modify the request
		c.Request.URL.Scheme = targetURL.Scheme
		c.Request.URL.Host = targetURL.Host
		c.Request.URL.Path = c.Param("path")

		// Let the reverse proxy do its job
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
func createReverseProxyBiz(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the target URL
		targetURL, _ := url.Parse(target)

		// Create the reverse proxy
		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		if c.Param("path") == "get_users" {
			conn, err := grpc.Dial(addressBiz, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			c2 := pbs.NewGetUsersClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			userId, _ := c.Get("userId")
			idx, _ := strconv.Atoi(userId.(string))
			idd := int32(idx)
			messageId, _ := c.Get("userId")
			messageid, _ := strconv.Atoi(messageId.(string))
			message_id := int32(messageid)
			authKey := c.GetHeader("userId")
			authkey, _ := strconv.Atoi(authKey)
			authkey32 := int32(authkey)

			req := &pbs.UserRequest{UserId: idd, MessageId: message_id, AuthKey: authkey32}
			r, err := c2.GetUsers(ctx, req)
			if err != nil {
				log.Fatalf("could not get users:  %v", err)

			}
			var x []pbs.User
			for _, user := range r.Users {
				x = append(x, *user)
			}

			c.JSON(http.StatusAccepted, gin.H{
				"messageId": r.MessageId, "users": x,
			})
			return

		}
		if c.Param("path") == "get_users_with_sql_inject" {
			conn, err := grpc.Dial(addressBiz, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			c2 := pbs.NewGetUsersClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			userId, _ := c.Get("userId")
			messageId, _ := c.Get("userId")
			messageid, _ := strconv.Atoi(messageId.(string))
			message_id := int32(messageid)
			authKey := c.GetHeader("userId")
			authkey, _ := strconv.Atoi(authKey)
			authkey32 := int32(authkey)

			req := &pbs.UserRequestWithSqlInject{UserId: userId.(string), MessageId: message_id, AuthKey: authkey32}
			r, err := c2.GetUsersWithSqlInject(ctx, req)
			if err != nil {
				log.Fatalf("could not get users:  %v", err)

			}
			var x []pbs.User
			for _, user := range r.Users {
				x = append(x, *user)
			}

			c.JSON(http.StatusAccepted, gin.H{
				"messageId": r.MessageId, "users": x,
			})
			return

		}

		// Modify the request
		c.Request.URL.Scheme = targetURL.Scheme
		c.Request.URL.Host = targetURL.Host
		c.Request.URL.Path = c.Param("path")
		authtoken := c.GetHeader("token")
		if authtoken == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "AUTH require",
			})
			return
			//send auth token to auth server and check token validity
		}

		proxy.ServeHTTP(c.Writer, c.Request)

	}
}
func main() {
	startServer(createRouter())
}
