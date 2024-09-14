package main

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	c := vpn.NewVPNClient(conn)
	dp := &vpn.DataPacket{
		TunnelId:         "192.88.99.2",
		EncryptedPayload: []byte("Hello, world!"),
	}

	tdc, err := c.TransferData(context.Background())
	if err != nil {
		panic(err)
	}

	for {
		tdc.Send(dp)
		resp, err := tdc.Recv()
		if err != nil {
			panic(err)
		}

		println("Received: ", string(resp.EncryptedPayload))
		time.Sleep(1 * time.Second)
	}
}
