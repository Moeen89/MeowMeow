package main

import (
	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/redis/go-redis/v9"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
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
