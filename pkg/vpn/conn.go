package vpn

import (
	"github.com/songgao/water"

	"github.com/yakumioto/emptiness/protos/vpn"
)

type Conn interface {
	Reading()
	Writing()
}

type connClient struct {
	tunnelID string // client tunnel id
	stream   vpn.VPN_TransferDataServer

	toServer chan *vpn.DataPacket // Send to VPN server
	out      chan *vpn.DataPacket // Send to client conn
}

func (c *connClient) Reading() {
	//TODO implement me
	panic("implement me")
}

func (c *connClient) Writing() {
	//TODO implement me
	panic("implement me")
}

type connTUNDevice struct {
	tunnelID string
	tun      *water.Interface

	toServer chan *vpn.DataPacket // Send to VPN server
	out      chan *vpn.DataPacket // Send to device conn
}

func (d *connTUNDevice) Reading() {
	//TODO implement me
	panic("implement me")
}

func (d *connTUNDevice) Writing() {
	//TODO implement me
	panic("implement me")
}
