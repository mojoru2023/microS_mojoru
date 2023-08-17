package main

import (
	"context"
	"fmt"
	proto "micro_s/mojoru/proto" //注意这里：修改成你自己的

	"go-micro.dev/v4"
	"go-micro.dev/v4/server"
)

// 定义结构体 作为方法调用方
type Greeter struct{}

//实现 .pb.micro.go中的Hello方法 定义rsp的返回值

func (*Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	fmt.Println("request:", req.Name)

	rsp.Greeting = "Hello" + req.Name
	return nil
}

func main() {
	//定义服务
	rpcServer := server.NewServer(
		server.Name("micro_s/mojoru.service"), //微服务的名称必须这么写
		server.Address(":8099"),
	)

	// Register Handlers  // 注册handler
	err := proto.RegisterGreeterHandler(rpcServer, &Greeter{})
	if err != nil {
		return
	}

	service := micro.NewService(
		micro.Server(rpcServer),
	)

	//服务初始化 	// optionally setup command line usage

	service.Init()

	//启动服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
