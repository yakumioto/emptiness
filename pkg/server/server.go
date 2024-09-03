package server

import (
	"context"

	"google.golang.org/protobuf/proto"

	"github.com/yakumioto/emptiness/pkg/crypto"
	"github.com/yakumioto/emptiness/pkg/tunnel"
	"github.com/yakumioto/emptiness/pkg/vpn"
)

type Server struct {
	vpn.UnimplementedVPNServer

	cryptoProvider crypto.Provider
	tunnelManager  *tunnel.Manager
}

func (s *Server) Connect(ctx context.Context, request *vpn.Request) (*vpn.Response, error) {
	payload, err := s.cryptoProvider.Decrypt(request.EncryptedPayload)
	if err != nil {
		return &vpn.Response{
			Code: vpn.StatusCode_INVALID_AUTHORIZATION,
		}, err
	}

	connectRequest := new(vpn.ConnectRequest)
	if err := proto.Unmarshal(payload, connectRequest); err != nil {
		return &vpn.Response{
			Code: vpn.StatusCode_UNKNOWN_ERROR,
		}, err
	}

	if s.tunnelManager.HasTunnel(connectRequest.TunnelId) {
		return &vpn.Response{
			Code: vpn.StatusCode_EXISTED_TUNNEL,
		}, nil
	}

	s.tunnelManager.AddTunnel(connectRequest.TunnelId, nil)

	return &vpn.Response{
		Code: vpn.StatusCode_OK,
	}, nil
}

func (s *Server) Disconnect(ctx context.Context, request *vpn.Request) (*vpn.Response, error) {
	payload, err := s.cryptoProvider.Decrypt(request.EncryptedPayload)
	if err != nil {
		return &vpn.Response{
			Code: vpn.StatusCode_INVALID_AUTHORIZATION,
		}, err
	}

	disconnectRequest := new(vpn.DisconnectRequest)
	if err := proto.Unmarshal(payload, disconnectRequest); err != nil {
		return &vpn.Response{
			Code: vpn.StatusCode_UNKNOWN_ERROR,
		}, err
	}

	s.tunnelManager.DelTunnel(disconnectRequest.TunnelId)

	return &vpn.Response{
		Code: vpn.StatusCode_OK,
	}, nil
}

func (s *Server) RegisterRoute(ctx context.Context, request *vpn.Request) (*vpn.Response, error) {
	payload, err := s.cryptoProvider.Decrypt(request.EncryptedPayload)
	if err != nil {
		return &vpn.Response{
			Code: vpn.StatusCode_INVALID_AUTHORIZATION,
		}, err
	}

	registerRouteRequest := new(vpn.RouteRequest)
	if err := proto.Unmarshal(payload, registerRouteRequest); err != nil {
		return &vpn.Response{
			Code: vpn.StatusCode_UNKNOWN_ERROR,
		}, err
	}

	if !s.tunnelManager.HasTunnel(registerRouteRequest.TunnelId) {
		return &vpn.Response{
			Code: vpn.StatusCode_NOT_FOUND_TUNNEL,
		}, nil
	}

	for _, ip := range registerRouteRequest.Route {
		s.tunnelManager.AddIP(registerRouteRequest.TunnelId, ip)
	}

	return &vpn.Response{
		Code: vpn.StatusCode_OK,
	}, nil
}

func (s *Server) UnregisterRoute(ctx context.Context, request *vpn.Request) (*vpn.Response, error) {
	payload, err := s.cryptoProvider.Decrypt(request.EncryptedPayload)
	if err != nil {
		return &vpn.Response{
			Code: vpn.StatusCode_INVALID_AUTHORIZATION,
		}, err
	}

	unregisterRouteRequest := new(vpn.RouteRequest)
	if err := proto.Unmarshal(payload, unregisterRouteRequest); err != nil {
		return &vpn.Response{
			Code: vpn.StatusCode_UNKNOWN_ERROR,
		}, err
	}

	if !s.tunnelManager.HasTunnel(unregisterRouteRequest.TunnelId) {
		return &vpn.Response{
			Code: vpn.StatusCode_NOT_FOUND_TUNNEL,
		}, nil
	}

	for _, ip := range unregisterRouteRequest.Route {
		s.tunnelManager.DelIP(unregisterRouteRequest.TunnelId, ip)
	}

	return &vpn.Response{
		Code: vpn.StatusCode_OK,
	}, nil
}

func (s *Server) TransferData(server vpn.VPN_TransferDataServer) error {
	for {
		dataPacket, err := server.Recv()
		if err != nil {
			return err
		}

	}
}
