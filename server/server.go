package main

import (
	"context"
	"fmt"
	"log"

	msg "../proto"
	"github.com/micro/go-micro"
)

// 聲明結構體
type Hello struct {
}

// 實現接口方法
func (h *Hello) Info(ctx context.Context, req *msg.InfoRequest, rep *msg.InfoResponse) error {
	rep.Msg = "hi " + req.Username
	return nil
}

func main() {
	// 1.得到微服務實例
	service := micro.NewService(
		// 設置微服務的名字，用來作訪問用的
		micro.Name("hello"),
	)

	// 2.初始化
	service.Init()

	// 3.服務註冊
	err := msg.RegisterHelloHandler(service.Server(), new(Hello))
	if err != nil {
		fmt.Println(err)
	}

	// 4.啟動微服務
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
