package server

import (
	"context"
	"fmt"

	pb "github.com/Pham-Xuan-Thanh/hdwallet/proto/btchdwallet"
	"github.com/Pham-Xuan-Thanh/hdwallet/wallet"
)

type Server struct {
	pb.UnimplementedWalletServer
}

func (s *Server) CreateWallet(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	fmt.Println()
	fmt.Println("\nCreating new wallet")

	wallet := wallet.CreateWallet()

	return wallet, nil
}

func (s *Server) GetWallet(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	fmt.Println()
	fmt.Println("\nGetting wallet data")

	wallet := wallet.DecodeWallet(in.Mnemonic)

	return wallet, nil
}

func (s *Server) GetBalance(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	fmt.Println()
	fmt.Println("\nGetting Balance data")

	balance := wallet.GetBalance(in.Address)

	return balance, nil
}
