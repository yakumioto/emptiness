package vpn

import (
	"net"
	"sync"

	"github.com/yakumioto/emptiness/protos/vpn"
)

// TunnelManager manages VPN tunnels, their associated connections and IPs
type TunnelManager struct {
	conns sync.Map // map[string]Conn
	ips   sync.Map // map[string]map[string]struct{}
}

// NewTunnelManager creates a new TunnelManager instance
func NewTunnelManager() *TunnelManager {
	return &TunnelManager{}
}

// AddTunnel adds a new tunnel with its associated connection
func (m *TunnelManager) AddTunnel(tunnelID string, conn Conn) {
	m.conns.Store(tunnelID, conn)
	m.ips.Store(tunnelID, make(map[string]struct{}))
}

// RemoveTunnel removes a tunnel and its associated data
func (m *TunnelManager) RemoveTunnel(tunnelID string) {
	m.conns.Delete(tunnelID)
	m.ips.Delete(tunnelID)
}

// AddIP adds an IP to a specific tunnel
func (m *TunnelManager) AddIP(tunnelID, ip string) {
	if ips, ok := m.ips.Load(tunnelID); ok {
		ips.(map[string]struct{})[ip] = struct{}{}
	}
}

// DelIP removes an IP from a specific tunnel
func (m *TunnelManager) DelIP(tunnelID, ip string) {
	if ips, ok := m.ips.Load(tunnelID); ok {
		delete(ips.(map[string]struct{}), ip)
	}
}

// HasIP checks if a specific tunnel has an IP
func (m *TunnelManager) HasIP(tunnelID, ip string) bool {
	if ips, ok := m.ips.Load(tunnelID); ok {
		_, exists := ips.(map[string]struct{})[ip]
		return exists
	}
	return false
}

// GetConn retrieves the connection associated with a tunnel
func (m *TunnelManager) GetConn(tunnelID string) (Conn, bool) {
	conn, ok := m.conns.Load(tunnelID)
	if !ok {
		return nil, false
	}
	return conn.(Conn), true
}

// Server represents the VPN server
type Server struct {
	netMask net.IPMask

	Tunnels *TunnelManager

	in chan *vpn.DataPacket // Receives packets from tun device and clients for processing
}
