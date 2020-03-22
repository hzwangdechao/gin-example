package main

import (
	"fmt"
	"gin-example/pkg/setting"
	"gin-example/routers"
	"net/http"
)

func main() {
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

	s.ListenAndServe()
	
}
