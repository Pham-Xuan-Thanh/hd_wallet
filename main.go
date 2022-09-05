package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Pham-Xuan-Thanh/hdwallet/config"
	pb "github.com/Pham-Xuan-Thanh/hdwallet/proto/btchdwallet"
	"github.com/Pham-Xuan-Thanh/hdwallet/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var conf = config.ParseConfig()

func main() {
	listenner, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.Port))
	if err != nil {
		log.Fatal("Fail to listen: %v ", err)
	}

	s := grpc.NewServer()

	pb.RegisterWalletServer(s, &server.Server{})
	reflection.Register(s)

	fmt.Printf("Server running at port: %s\n", conf.Port)

	if err := s.Serve(listenner); err != nil {
		log.Fatal("Fail to establish server: %v", err)
	}
}
