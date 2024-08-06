package vpn

import (
	"net"
	"sync"

	"github.com/songgao/water"

	"github.com/yakumioto/emptiness/protos/vpn"
)

type clientConn struct {
	tunnelID string // client tunnel id

	ip  net.IP   // client vpn ip
	ips sync.Map // client child vpn ip list

	stream vpn.VPN_TransferDataServer

	in  chan *vpn.DataPacket // Send to VPN server
	out chan *vpn.DataPacket // Send to client conn
}

func (c *clientConn) AddIP(ip string) {
	c.ips.Store(ip, true)
}

func (c *clientConn) DelIP(ip string) {
	c.ips.Delete(ip)
}

func (c *clientConn) HasIP(ip string) bool {
	_, ok := c.ips.Load(ip)
	return ok
}

type clientPool struct {
	clients sync.Map
}

func (c *clientPool) AddClient(ip string, client *clientConn) {
	client.AddIP(ip)
	c.clients.Store(ip, client)
}

func (c *clientPool) GetClient(ip string) *clientConn {
	if c, ok := c.clients.Load(ip); ok {
		return c.(*clientConn)
	}

	return nil
}

func (c *clientPool) DelClient(ip string) {
	client := c.GetClient(ip)
	if client != nil {
		client.DelIP(ip)
	}

	c.clients.Delete(ip)
}

type Server struct {
	tun     *water.Interface
	tunName string

	ip      net.IP
	netMask net.IPMask

	clients *clientPool

	in  chan *vpn.DataPacket // Recv by tun device and client
	out chan *vpn.DataPacket // Send to tun device
}
