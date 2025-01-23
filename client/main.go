package main

import (
	"context"
	"flag"
	pb "grpc-demo/rpc/hello" // Protocol Buffers
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", "hello world", "Name to greet")
)

func main() {
	flag.Parse()
	// 1. 创建连接
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 2. 创建客户端
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second) // 设置超时时间
	defer cancel()                                                        // 调用 cancel 函数释放资源
	// 3. 调用服务
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	// 4. 打印返回结果
	log.Printf("Greeting: %s", r.GetMessage())
}
