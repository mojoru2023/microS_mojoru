package main

import (
	"bufio"
	"context"
	"fmt"
	proto "micro_s/mojoru/proto" //注意这里：修改成你自己的

	"os"

	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
)

func main() {
	//创建一个新的服务 命名

	service := micro.NewService(
		micro.Client(client.NewClient()),
	)

	//服务初始化
	service.Init()
	//创建服务 绑定客户端 这个方法是在proto生成的文件中定义的
	client := proto.NewGreeterService("micro_s/mojoru.service", service.Client()) //微服务的名称必须这么写

	//调用Hello方法 Hello方法同样是在proto生成的文件中定义的
	rsp, err := client.Hello(context.TODO(), &proto.Request{Name: "World"})
	if err != nil {
		fmt.Println(err)
	}

	//打印结果
	fmt.Println(rsp.Greeting)
	fmt.Println("Press Enter key to exit the program...")
	in := bufio.NewReader(os.Stdin)
	_, _, _ = in.ReadLine()

	// service := micro.NewService(
	// 	micro.Client(client.NewClient()),
	// )

	// service.Init()
	// client := proto.NewGreeterService("micro_s/mojoru.service", service.Client())

	// rsp, err := client.Hello(context.TODO(), &proto.Request{Name: "BOSSMA"})
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(rsp)

	// fmt.Println("Press Enter key to exit the program...")
	// in := bufio.NewReader(os.Stdin)
	// _, _, _ = in.ReadLine()
}
