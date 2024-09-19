package tunnel

import (
	"sync"

	pb "github.com/yakumioto/emptiness/protobuf"
)

// Manager manages VPN tunnels, their associated connections and IPs
type Manager struct {
	tunnelMap sync.Map // map[string]Tunneler
	routeMap  sync.Map // map[string]map[string]struct{}

	in chan *pb.DataPacket // Read by vpn server transfer data to tunnel.
}

// NewManager creates a new Manager instance
func NewManager() *Manager {
	return &Manager{
		tunnelMap: sync.Map{},
		routeMap:  sync.Map{},
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
		m.routeMap.Store(tunnelID, make(map[string]struct{}))
	}
}

// DelTunnel removes a tunnel and its associated data
func (m *Manager) DelTunnel(tunnelID string) {
	m.tunnelMap.Delete(tunnelID)
	m.routeMap.Delete(tunnelID)
}

// AddRoute adds an IP to a specific tunnel
func (m *Manager) AddRoute(tunnelID, ip string) {
	if !m.hasRoute(tunnelID, ip) {
		if ips, ok := m.routeMap.Load(tunnelID); ok {
			ips.(map[string]struct{})[ip] = struct{}{}
		}
	}
}

// DelRoute removes an IP from a specific tunnel
func (m *Manager) DelRoute(tunnelID, ip string) {
	if ips, ok := m.routeMap.Load(tunnelID); ok {
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

// hasRoute checks if a specific tunnel has an IP
func (m *Manager) hasRoute(tunnelID, ip string) bool {
	if ips, ok := m.routeMap.Load(tunnelID); ok {
		_, exists := ips.(map[string]struct{})[ip]
		return exists
	}
	return false
}

// hasRouteInAnyTunnel checks if an IP is in any tunnel
func (m *Manager) hasRouteInAnyTunnel(ip string) (exists bool) {
	m.routeMap.Range(func(key, _ interface{}) bool {
		if m.hasRoute(key.(string), ip) {
			exists = true
			return false
		}
		return true
	})

	return
}
