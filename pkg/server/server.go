package server

import (
	"context"
	"log"

	"google.golang.org/protobuf/proto"

	"github.com/yakumioto/emptiness/pkg/crypto"
	"github.com/yakumioto/emptiness/pkg/tunnel"
	pb "github.com/yakumioto/emptiness/protobuf"
)

type Server struct {
	pb.UnimplementedVPNServer

	CryptoProvider crypto.Provider
	TunnelManager  *tunnel.Manager
}

func (s *Server) Connect(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	payload, err := s.CryptoProvider.Decrypt(request.EncryptedPayload)
	if err != nil {
		return &pb.Response{
			StatusCode: pb.StatusCode_INVALID_AUTHORIZATION,
		}, err
	}

	connectRequest := new(pb.ConnectRequest)
	if err := proto.Unmarshal(payload, connectRequest); err != nil {
		return &pb.Response{
			StatusCode: pb.StatusCode_UNKNOWN_ERROR,
		}, err
	}

	log.Println("Connect request received: ", connectRequest)

	return &pb.Response{
		StatusCode: pb.StatusCode_OK,
	}, nil
}

func (s *Server) Disconnect(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	payload, err := s.CryptoProvider.Decrypt(request.EncryptedPayload)
	if err != nil {
		return &pb.Response{
			StatusCode: pb.StatusCode_INVALID_AUTHORIZATION,
		}, err
	}

	disconnectRequest := new(pb.DisconnectRequest)
	if err := proto.Unmarshal(payload, disconnectRequest); err != nil {
		return &pb.Response{
			StatusCode: pb.StatusCode_UNKNOWN_ERROR,
		}, err
	}

	s.TunnelManager.DelTunnel(disconnectRequest.TunnelId)

	return &pb.Response{
		StatusCode: pb.StatusCode_OK,
	}, nil
}

func (s *Server) RegisterRoute(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	payload, err := s.CryptoProvider.Decrypt(request.EncryptedPayload)
	if err != nil {
		return &pb.Response{
			StatusCode: pb.StatusCode_INVALID_AUTHORIZATION,
		}, err
	}

	registerRouteRequest := new(pb.RouteRequest)
	if err := proto.Unmarshal(payload, registerRouteRequest); err != nil {
		return &pb.Response{
			StatusCode: pb.StatusCode_UNKNOWN_ERROR,
		}, err
	}

	for _, ip := range registerRouteRequest.Route {
		s.TunnelManager.AddIP(registerRouteRequest.TunnelId, ip)
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

	for _, ip := range unregisterRouteRequest.Route {
		s.TunnelManager.DelIP(unregisterRouteRequest.TunnelId, ip)
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
