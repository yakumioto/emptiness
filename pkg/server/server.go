package server

import (
	"context"

	"github.com/yakumioto/emptiness/pkg/vpn"
)

type Server struct {
	vpn.UnimplementedVPNServer
}

func (s *Server) Connect(ctx context.Context, request *vpn.ConnectRequest) (*vpn.ConnectResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Server) Disconnect(ctx context.Context, request *vpn.DisconnectRequest) (*vpn.DisconnectResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Server) RegisterRoute(ctx context.Context, request *vpn.RouteRequest) (*vpn.RouteResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Server) UnregisterRoute(ctx context.Context, request *vpn.RouteRequest) (*vpn.RouteResponse, error) {
	// TODO implement me
	panic("implement me")
}
