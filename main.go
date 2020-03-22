package main

import (
	"context"
	"fmt"
	"gin-example/pkg/setting"
	"gin-example/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//endless.DefaultReadTimeOut = setting.ReadTimeout
	//endless.DefaultWriteTimeOut = setting.WriteTimeOut
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	//
	//server := endless.NewServer(endPoint, routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err :%v", err)
	//
	//}
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		TLSConfig:      nil,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeOut,
		IdleTimeout:    0,
		MaxHeaderBytes: 1 << 20,
	}
	//s.ListenAndServe()

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s \n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("shutdown server ....")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown:%s", err)
	}
	log.Println("Server exiting")

}
