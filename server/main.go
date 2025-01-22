package main

import (
	"context"
	"log"
	"net"

	pb "github.com/wcx0206/grpc-firstTry/rpc/hello" // 修改为正确的导入路径

	"google.golang.org/grpc"
)

// 定义服务结构体
type server struct {
	pb.UnimplementedGreeterServer
}

// 实现 SayHello 方法
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "Hello " + req.GetName(),
	}, nil
}

func main() {
	// 1. 创建 TCP 监听器
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 2. 创建 gRPC 服务器
	s := grpc.NewServer()

	// 3. 注册服务
	pb.RegisterGreeterServer(s, &server{})

	// 4. 启动服务
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
