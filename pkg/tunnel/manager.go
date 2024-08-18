package tunnel

import "sync"

// Manager manages VPN tunnels, their associated connections and IPs
type Manager struct {
	connMap sync.Map // map[string]Tunneler
	ipMap   sync.Map // map[string]map[string]struct{}
}

// NewTunnelManager creates a new Manager instance
func NewTunnelManager() *Manager {
	return &Manager{}
}

// AddTunnel adds a new tunnel with its associated connection
func (m *Manager) AddTunnel(tunnelID string, conn Tunnel) {
	m.connMap.Store(tunnelID, conn)
	m.ipMap.Store(tunnelID, make(map[string]struct{}))
}

// RemoveTunnel removes a tunnel and its associated data
func (m *Manager) RemoveTunnel(tunnelID string) {
	m.connMap.Delete(tunnelID)
	m.ipMap.Delete(tunnelID)
}

// GetConn retrieves the connection associated with a tunnel
func (m *Manager) GetConn(tunnelID string) (Tunnel, bool) {
	conn, ok := m.connMap.Load(tunnelID)
	if !ok {
		return nil, false
	}
	return conn.(Tunnel), true
}

// AddIP adds an IP to a specific tunnel
func (m *Manager) AddIP(tunnelID, ip string) {
	if ips, ok := m.ipMap.Load(tunnelID); ok {
		ips.(map[string]struct{})[ip] = struct{}{}
	}
}

// DelIP removes an IP from a specific tunnel
func (m *Manager) DelIP(tunnelID, ip string) {
	if ips, ok := m.ipMap.Load(tunnelID); ok {
		delete(ips.(map[string]struct{}), ip)
	}
}

// HasIP checks if a specific tunnel has an IP
func (m *Manager) HasIP(tunnelID, ip string) bool {
	if ips, ok := m.ipMap.Load(tunnelID); ok {
		_, exists := ips.(map[string]struct{})[ip]
		return exists
	}
	return false
}

// HasIPInAnyTunnel checks if an IP is in any tunnel
func (m *Manager) HasIPInAnyTunnel(ip string) bool {
	var exists bool
	m.ipMap.Range(func(_, ips interface{}) bool {
		if _, ok := ips.(map[string]struct{})[ip]; ok {
			exists = true
			return false
		}
		return true
	})
	return exists
}
