package main

import (
	"bookstore/store/factory"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s, err := factory.New("memo") // 创建图书存储模块实例
	if err != nil {
		panic(err)
	}

	srv := server.NewBookStoreServer(":8080", s) // 创建http服务实例

	errChan, err := srv.ListenAndServe() //运行http服务
	if err != nil {
		log.Println("web server start failed:", err)
		return
	}
	log.Println("web server start ok")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select { //监听来自errChan以及c的事件
	case err = <-errChan:
		log.Println("web server run failed:", err)
		ctx, cf := context.WithTimeout(context.Background(), time.Second)
		defer cf()
		err = srv.Shutdown(ctx) //优雅关闭http服务
	}
	if err != nil {
		log.Println("bookstore program exit error", err)
		return
	}
	log.Println("bookstore program exit ok")
}
