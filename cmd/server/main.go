package main

import (
	"net"

	"google.golang.org/grpc"

	"github.com/yakumioto/emptiness/crypto"
	"github.com/yakumioto/emptiness/server"
	"github.com/yakumioto/emptiness/tunnel"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	vpn.RegisterVPNServer(s, &server.Server{
		CryptoProvider: &crypto.None{},
		TunnelManager:  tunnel.NewManager(),
	})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
