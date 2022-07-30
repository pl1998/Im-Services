/**
  @author:panliang
  @data:2022/7/17
  @note
**/
package server

import (
	"google.golang.org/grpc"
	grpcMessage "im-services/server/grpc/message"
	"log"
	"net"
)

var RpcServer = grpc.NewServer()

func StartGrpc() {

	var message grpcMessage.ImGrpcMessage

	grpcMessage.RegisterImMessageServer(RpcServer, message)

	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("grpc服务启动失败", err)
	}
	_ = RpcServer.Serve(listener)

}
