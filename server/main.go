package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"server/inits"
	"server/public"
	"server/routes"
	"time"
)

func main() {
	app := gin.Default()
	err := inits.InjectData() //初始化数据
	if err != nil {
		log.Println("init data err:", err.Error())
	}
	s := http.Server{
		Addr:           public.SEVER_ADDRESS,
		Handler:        app,
		ReadTimeout:    time.Duration(public.OUTTIME_SECONDS) * time.Second,
		WriteTimeout:   time.Duration(public.OUTTIME_SECONDS) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	routes.SetupRoutes(app) //注册路由
	log.Println("server run at port", public.SEVER_ADDRESS)
	log.Fatalln("start server err:", s.ListenAndServe())

}
