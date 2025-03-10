package net_probe

import (
	"fmt"
	"net"

	"github.com/csmali/bettercap/packets"
)

func (mod *Prober) sendProbeNBNS(from net.IP, from_hw net.HardwareAddr, ip net.IP) {
	name := fmt.Sprintf("%s:%d", ip, packets.NBNSPort)
	if addr, err := net.ResolveUDPAddr("udp", name); err != nil {
		mod.Debug("could not resolve %s.", name)
	} else if con, err := net.DialUDP("udp", nil, addr); err != nil {
		mod.Debug("could not dial %s.", name)
	} else {
		defer con.Close()
		if wrote, _ := con.Write(packets.NBNSRequest); wrote > 0 {
			mod.Session.Queue.TrackSent(uint64(wrote))
		} else {
			mod.Session.Queue.TrackError()
		}
	}
}
