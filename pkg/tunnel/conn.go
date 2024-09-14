package tunnel

import pb "github.com/yakumioto/emptiness/protobuf"

type conn struct {
	tunnelID string // client tunnel id
	stream   pb.VPN_TransferDataServer

	out chan *pb.DataPacket // Read by gRPC stream out to tunnel manager.
	in  chan *pb.DataPacket // Read by tunnel manager out to gRPC stream.
}

func (c *conn) Reading() {
	// TODO implement me
	panic("implement me")
}

func (c *conn) Writing() {
	// TODO implement me
	panic("implement me")
}
