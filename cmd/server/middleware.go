package main

import (
	"awesomeProject/internal/server"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(serv *server.Server) func(c *gin.Context) {
	return cors.New(cors.Config{
		AllowOrigins:     serv.Cors.AllowedOrigin,
		AllowMethods:     serv.Cors.AllowedMethods,
		AllowHeaders:     serv.Cors.AllowedHeaders,
		AllowCredentials: serv.Cors.AllowedCredentials,
	})
}

func BasicAuthRequired(serv *server.Server) func(c *gin.Context) {

	return func(c *gin.Context) {
		user, password, hasAuth := c.Request.BasicAuth()
		if hasAuth && user == serv.InternalAuth.Username && password == serv.InternalAuth.Password {
			c.Next()
		} else {
			c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
}

type handler func(ctx *gin.Context)

func wrapHandler(h handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(c)
	}
}
