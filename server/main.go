package main

import (
	"context"
	"flag"
	"fmt"
	pb "grpc-demo/rpc/hello"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "服务器端口")
)

// 定义服务结构体
type server struct {
	pb.UnimplementedGreeterServer // 保持实现接口的兼容性
}

// 实现 SayHello 方法
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("收到请求: %v", in.GetName())
	return &pb.HelloReply{
		Message: "Hello " + in.GetName(),
	}, nil
}

func main() {
	// 解析命令行参数
	flag.Parse()

	// 监听端口
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	// 创建 gRPC 服务
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("服务器监听端口: %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
