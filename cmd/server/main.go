package main

import (
	"awesomeProject/cmd/common"
	"flag"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	configFilePtr := flag.String("config-file", "./configs/dev.toml", "Config file to use")
	config := common.NewConfig(*configFilePtr)
	engine := gin.New()
	server := common.MustCreateServer(config)
	AddKnockKnock(engine, server)
	engine.Use(CORSMiddleware(server))

	s := &http.Server{
		Addr:           ":" + strconv.FormatUint(config.Server.Port, 10),
		Handler:        engine,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   20 * time.Second,
		IdleTimeout:    65 * time.Second,
		MaxHeaderBytes: 5 * http.DefaultMaxHeaderBytes,
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
