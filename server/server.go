package server

import (
	"context"
	"log"

	"google.golang.org/protobuf/proto"

	"github.com/yakumioto/emptiness/crypto"
	pb "github.com/yakumioto/emptiness/protobuf"
	"github.com/yakumioto/emptiness/tunnel"
)

type Server struct {
	pb.UnimplementedVPNServer

	CryptoProvider crypto.Provider
	TunnelManager  *tunnel.Manager
}

func (s *Server) RegisterRoute(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	payload, err := s.CryptoProvider.Decrypt(request.EncryptedPayload)
	if err != nil {
		return &pb.Response{
			StatusCode: pb.StatusCode_INVALID_AUTHORIZATION,
		}, err
	}

	registerRouteRequest := new(pb.RouteRequest)
	if err = proto.Unmarshal(payload, registerRouteRequest); err != nil {
		return &pb.Response{
			StatusCode: pb.StatusCode_UNKNOWN_ERROR,
		}, err
	}

	for _, route := range registerRouteRequest.Routes {
		s.TunnelManager.AddRoute(registerRouteRequest.TunnelId, route)
	}

	return &pb.Response{
		StatusCode: pb.StatusCode_OK,
	}, nil
}

func (s *Server) UnregisterRoute(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	payload, err := s.CryptoProvider.Decrypt(request.EncryptedPayload)
	if err != nil {
		return &pb.Response{
			StatusCode: pb.StatusCode_INVALID_AUTHORIZATION,
		}, err
	}

	unregisterRouteRequest := new(pb.RouteRequest)
	if err := proto.Unmarshal(payload, unregisterRouteRequest); err != nil {
		return &pb.Response{
			StatusCode: pb.StatusCode_UNKNOWN_ERROR,
		}, err
	}

	for _, route := range unregisterRouteRequest.Routes {
		s.TunnelManager.DelRoute(unregisterRouteRequest.TunnelId, route)
	}

	return &pb.Response{
		StatusCode: pb.StatusCode_OK,
	}, nil
}

func (s *Server) TransferData(server pb.VPN_TransferDataServer) error {
	for {
		request, err := server.Recv()
		if err != nil {
			return err
		}

		payload, err := s.CryptoProvider.Decrypt(request.EncryptedPayload)
		if err != nil {
			return err
		}

		dataPacket := new(pb.DataPacket)
		if err := proto.Unmarshal(payload, dataPacket); err != nil {
			return err
		}

		log.Println("Data packet received: ", string(payload))

		server.Send(&pb.StreamResponse{})
	}
}
