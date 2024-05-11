package main

import (
	"fmt"
	"ginDemo/pkg/setting"
	"ginDemo/routers"
	"net/http"
)

func main() {
	r := routers.InitGinServer()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
	//r.Run() // listen and serve on 0.0.0.0:8080
}
