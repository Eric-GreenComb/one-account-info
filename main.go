package main

import (
	"log"
	"net/http"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/golang/sync/errgroup"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Eric-GreenComb/one-account-info/config"
	"github.com/Eric-GreenComb/one-account-info/ethereum"
	"github.com/Eric-GreenComb/one-account-info/handler"
	"github.com/Eric-GreenComb/one-account-info/persist"
)

var (
	g errgroup.Group
)

func main() {
	if config.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	persist.InitDatabase()

	ethereum.Init()

	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.MaxMultipartMemory = 64 << 20 // 64 MiB

	router.Use(Cors())

	/* api base */
	r0 := router.Group("/")
	{
		r0.GET("", handler.Index)
		r0.GET("health", handler.Health)
	}

	r1 := router.Group("/acount")
	{
		r1.POST("/create", handler.CreateAccount)
		r1.POST("/update", handler.UpdateAccount)
		r1.GET("/info/:code", handler.GetAccountInfo)
		r1.GET("/list", handler.ListAccount)
	}

	r2 := router.Group("/balance")
	{
		r2.GET("/eth/:code", handler.GetEtherBalance)
		r2.GET("/token/:code", handler.GetTokenBalance)
	}

	for _, _port := range config.Server.Port {
		server := &http.Server{
			Addr:         ":" + _port,
			Handler:      router,
			ReadTimeout:  300 * time.Second,
			WriteTimeout: 300 * time.Second,
		}

		g.Go(func() error {
			return gracehttp.Serve(server)
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
