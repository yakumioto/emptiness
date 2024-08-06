package vpn

import "github.com/songgao/water"

type Client struct {
	tun     *water.Interface
	tunName string
}
