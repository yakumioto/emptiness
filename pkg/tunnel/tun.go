package tunnel

import (
	"github.com/songgao/water"

	pb "github.com/yakumioto/emptiness/protobuf"
)

type tun struct {
	tunnelID string
	tun      *water.Interface

	out chan *pb.DataPacket // Read by device tun out to tunnel manager.
	in  chan *pb.DataPacket // Read by tunnel manager out to device tun.
}

func (d *tun) Reading() {
	// TODO implement me
	panic("implement me")
}

func (d *tun) Writing() {
	// TODO implement me
	panic("implement me")
}
