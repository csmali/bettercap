package modules

import (
	"github.com/csmali/bettercap/modules/any_proxy"
	"github.com/csmali/bettercap/modules/api_rest"
	"github.com/csmali/bettercap/modules/arp_spoof"
	"github.com/csmali/bettercap/modules/ble"
	"github.com/csmali/bettercap/modules/c2"
	"github.com/csmali/bettercap/modules/caplets"
	"github.com/csmali/bettercap/modules/dhcp6_spoof"
	"github.com/csmali/bettercap/modules/dns_spoof"
	"github.com/csmali/bettercap/modules/events_stream"
	"github.com/csmali/bettercap/modules/gps"
	"github.com/csmali/bettercap/modules/hid"
	"github.com/csmali/bettercap/modules/http_proxy"
	"github.com/csmali/bettercap/modules/http_server"
	"github.com/csmali/bettercap/modules/https_proxy"
	"github.com/csmali/bettercap/modules/https_server"
	"github.com/csmali/bettercap/modules/mac_changer"
	"github.com/csmali/bettercap/modules/mdns_server"
	"github.com/csmali/bettercap/modules/mysql_server"
	"github.com/csmali/bettercap/modules/net_probe"
	"github.com/csmali/bettercap/modules/net_recon"
	"github.com/csmali/bettercap/modules/net_sniff"
	"github.com/csmali/bettercap/modules/packet_proxy"
	"github.com/csmali/bettercap/modules/syn_scan"
	"github.com/csmali/bettercap/modules/tcp_proxy"
	"github.com/csmali/bettercap/modules/ticker"
	"github.com/csmali/bettercap/modules/ui"
	"github.com/csmali/bettercap/modules/update"
	"github.com/csmali/bettercap/modules/wifi"
	"github.com/csmali/bettercap/modules/wol"

	"github.com/csmali/bettercap/session"
)

func LoadModules(sess *session.Session) {
	sess.Register(any_proxy.NewAnyProxy(sess))
	sess.Register(arp_spoof.NewArpSpoofer(sess))
	sess.Register(api_rest.NewRestAPI(sess))
	sess.Register(ble.NewBLERecon(sess))
	sess.Register(dhcp6_spoof.NewDHCP6Spoofer(sess))
	sess.Register(net_recon.NewDiscovery(sess))
	sess.Register(dns_spoof.NewDNSSpoofer(sess))
	sess.Register(events_stream.NewEventsStream(sess))
	sess.Register(gps.NewGPS(sess))
	sess.Register(http_proxy.NewHttpProxy(sess))
	sess.Register(http_server.NewHttpServer(sess))
	sess.Register(https_proxy.NewHttpsProxy(sess))
	sess.Register(https_server.NewHttpsServer(sess))
	sess.Register(mac_changer.NewMacChanger(sess))
	sess.Register(mysql_server.NewMySQLServer(sess))
	sess.Register(mdns_server.NewMDNSServer(sess))
	sess.Register(net_sniff.NewSniffer(sess))
	sess.Register(packet_proxy.NewPacketProxy(sess))
	sess.Register(net_probe.NewProber(sess))
	sess.Register(syn_scan.NewSynScanner(sess))
	sess.Register(tcp_proxy.NewTcpProxy(sess))
	sess.Register(ticker.NewTicker(sess))
	sess.Register(wifi.NewWiFiModule(sess))
	sess.Register(wol.NewWOL(sess))
	sess.Register(hid.NewHIDRecon(sess))
	sess.Register(c2.NewC2(sess))

	sess.Register(caplets.NewCapletsModule(sess))
	sess.Register(update.NewUpdateModule(sess))
	sess.Register(ui.NewUIModule(sess))
}
