package main

import (
	loginendpoint "CoreSystemBase/service/api/login"
	//	middleware "CoreSystemBase/service/middleware"
	"flag"
	"fmt"
	"time"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
	"github.com/itsjamie/gin-cors"
)

var (
	port  = flag.Int64("p", 80, "port")
	debug = flag.Bool("d", false, "debug model")
)

func ready() {

	flag.Parse()
	if *port == 8060 {
		gin.SetMode(gin.ReleaseMode)
		log.Info(fmt.Sprintf("ReleaseMode"))
	}
}

func main() {
	ready()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET,POST,PUT,DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, Access-Control-Allow-Headers, LoginToken,X-Requested-With, X-CSRF-Token",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	authorized := r.Group("/user")
	{
		authorized.POST("login", loginendpoint.Login)
	}

	p := fmt.Sprintf(":%d", *port)
	log.Info("listen port " + p)
	r.Run(p)
}
