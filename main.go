package main

import (
	"fmt"
	"github.com/robfig/cron"
	"net/http"
	"os"
	"time"
	"urlnat/global"
	"urlnat/router"
	"urlnat/service"
)

func main() {
	if err := global.InitDB(); err != nil {
		panic(err)
	}
	c := cron.New()
	if err := c.AddFunc("0 0 0 * * *", func() {
		_ = service.ClearUrlNat()
	}); err != nil {
		panic(err)
	}
	c.Start()
	port := os.Getenv("PORT")
	engine := router.Router()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
