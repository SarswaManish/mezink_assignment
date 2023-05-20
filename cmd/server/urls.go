package main

import (
	"awesomeProject/internal/server"

	"github.com/gin-gonic/gin"
)

func AddKnockKnock(engine *gin.Engine, serv *server.Server) {
	engine.GET("/knockknock", wrapHandler(serv.KnockKnock))
}

func AddPrivate(private *gin.RouterGroup, serv *server.Server) {
	private.Use(BasicAuthRequired(serv))
	{
		private.GET("/fetch/records", wrapHandler(serv.FetchRecordsView))
	}
}
