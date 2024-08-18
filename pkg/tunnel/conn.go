package tunnel

import "github.com/yakumioto/emptiness/protos/vpn"

type conn struct {
	tunnelID string // client tunnel id
	stream   vpn.VPN_TransferDataServer

	out chan *vpn.DataPacket // Send to VPN server
	in  chan *vpn.DataPacket // Send to client conn
}

func (c *conn) Reading() {
	// TODO implement me
	panic("implement me")
}

func (c *conn) Writing() {
	// TODO implement me
	panic("implement me")
}
