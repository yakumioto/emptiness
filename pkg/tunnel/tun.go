package tunnel

import (
	"github.com/songgao/water"

	"github.com/yakumioto/emptiness/pkg/vpn"
)

type tun struct {
	tunnelID string
	tun      *water.Interface

	out chan *vpn.DataPacket // Send to VPN server
	in  chan *vpn.DataPacket // Send to device conn
}

func (d *tun) Reading() {
	// TODO implement me
	panic("implement me")
}

func (d *tun) Writing() {
	// TODO implement me
	panic("implement me")
}
