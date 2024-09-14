package tunnel

import (
	"sync"

	pb "github.com/yakumioto/emptiness/protobuf"
)

// Manager manages VPN tunnels, their associated connections and IPs
type Manager struct {
	tunnelMap sync.Map // map[string]Tunneler
	ipMap     sync.Map // map[string]map[string]struct{}

	in chan *pb.DataPacket // Read by vpn server transfer data to tunnel.
}

// NewManager creates a new Manager instance
func NewManager() *Manager {
	return &Manager{
		tunnelMap: sync.Map{},
		ipMap:     sync.Map{},
	}
}

func (m *Manager) Running() {
	// TODO implement me
	panic("implement me")
}

// AddTunnel adds a new tunnel with its associated connection
func (m *Manager) AddTunnel(tunnelID string, tunnel Tunnel) {
	if !m.hasTunnel(tunnelID) {
		m.tunnelMap.Store(tunnelID, tunnel)
		m.ipMap.Store(tunnelID, make(map[string]struct{}))
	}
}

// DelTunnel removes a tunnel and its associated data
func (m *Manager) DelTunnel(tunnelID string) {
	m.tunnelMap.Delete(tunnelID)
	m.ipMap.Delete(tunnelID)
}

// AddIP adds an IP to a specific tunnel
func (m *Manager) AddIP(tunnelID, ip string) {
	if !m.hasIPInAnyTunnel(ip) {
		if ips, ok := m.ipMap.Load(tunnelID); ok {
			ips.(map[string]struct{})[ip] = struct{}{}
		}
	}
}

// DelIP removes an IP from a specific tunnel
func (m *Manager) DelIP(tunnelID, ip string) {
	if ips, ok := m.ipMap.Load(tunnelID); ok {
		delete(ips.(map[string]struct{}), ip)
	}
}

// HasTunnel checks if a tunnel exists
func (m *Manager) hasTunnel(tunnelID string) bool {
	_, ok := m.getTunnel(tunnelID)
	return ok
}

// getTunnel retrieves the connection associated with a tunnel
func (m *Manager) getTunnel(tunnelID string) (Tunnel, bool) {
	conn, ok := m.tunnelMap.Load(tunnelID)
	if !ok {
		return nil, false
	}
	return conn.(Tunnel), true
}

// hasIP checks if a specific tunnel has an IP
func (m *Manager) hasIP(tunnelID, ip string) bool {
	if ips, ok := m.ipMap.Load(tunnelID); ok {
		_, exists := ips.(map[string]struct{})[ip]
		return exists
	}
	return false
}

// hasIPInAnyTunnel checks if an IP is in any tunnel
func (m *Manager) hasIPInAnyTunnel(ip string) (exists bool) {
	m.ipMap.Range(func(key, _ interface{}) bool {
		if m.hasIP(key.(string), ip) {
			exists = true
			return false
		}
		return true
	})

	return
}
