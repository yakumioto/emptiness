package vpn

import (
	"net"

	"github.com/yakumioto/emptiness/pkg/crypto"
	"github.com/yakumioto/emptiness/pkg/tunnel"
	pb "github.com/yakumioto/emptiness/protobuf"
)

type VPN struct {
	localIP        net.IP
	netMask        net.IPMask
	cryptoProvider crypto.Provider
	tunnelManager  *tunnel.Manager
	inPacket       chan *pb.DataPacket
}
