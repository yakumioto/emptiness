package vpn

import (
	"github.com/songgao/water"

	"github.com/yakumioto/emptiness/protos/vpn"
)

type Tunnel interface {
	Reading()
	Writing()
}

type clientTunnel struct {
	tunnelID string // client tunnel id
	stream   vpn.VPN_TransferDataServer

	out chan *vpn.DataPacket // Send to VPN server
	in  chan *vpn.DataPacket // Send to client conn
}

func (c *clientTunnel) Reading() {
	// TODO implement me
	panic("implement me")
}

func (c *clientTunnel) Writing() {
	// TODO implement me
	panic("implement me")
}

type tunTunnel struct {
	tunnelID string
	tun      *water.Interface

	out chan *vpn.DataPacket // Send to VPN server
	in  chan *vpn.DataPacket // Send to device conn
}

func (d *tunTunnel) Reading() {
	// TODO implement me
	panic("implement me")
}

func (d *tunTunnel) Writing() {
	// TODO implement me
	panic("implement me")
}
