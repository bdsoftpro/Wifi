package modules

import (
	"github.com/bdsoftpro/wifi/modules/any_proxy"
	"github.com/bdsoftpro/wifi/modules/api_rest"
	"github.com/bdsoftpro/wifi/modules/arp_spoof"
	"github.com/bdsoftpro/wifi/modules/ble"
	"github.com/bdsoftpro/wifi/modules/c2"
	"github.com/bdsoftpro/wifi/modules/caplets"
	"github.com/bdsoftpro/wifi/modules/dhcp6_spoof"
	"github.com/bdsoftpro/wifi/modules/dns_spoof"
	"github.com/bdsoftpro/wifi/modules/events_stream"
	"github.com/bdsoftpro/wifi/modules/gps"
	"github.com/bdsoftpro/wifi/modules/hid"
	"github.com/bdsoftpro/wifi/modules/http_proxy"
	"github.com/bdsoftpro/wifi/modules/http_server"
	"github.com/bdsoftpro/wifi/modules/https_proxy"
	"github.com/bdsoftpro/wifi/modules/https_server"
	"github.com/bdsoftpro/wifi/modules/mac_changer"
	"github.com/bdsoftpro/wifi/modules/mdns_server"
	"github.com/bdsoftpro/wifi/modules/mysql_server"
	"github.com/bdsoftpro/wifi/modules/ndp_spoof"
	"github.com/bdsoftpro/wifi/modules/net_probe"
	"github.com/bdsoftpro/wifi/modules/net_recon"
	"github.com/bdsoftpro/wifi/modules/net_sniff"
	"github.com/bdsoftpro/wifi/modules/packet_proxy"
	"github.com/bdsoftpro/wifi/modules/syn_scan"
	"github.com/bdsoftpro/wifi/modules/tcp_proxy"
	"github.com/bdsoftpro/wifi/modules/ticker"
	"github.com/bdsoftpro/wifi/modules/ui"
	"github.com/bdsoftpro/wifi/modules/update"
	"github.com/bdsoftpro/wifi/modules/wifi"
	"github.com/bdsoftpro/wifi/modules/wol"

	"github.com/bdsoftpro/wifi/session"
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
	sess.Register(ndp_spoof.NewNDPSpoofer(sess))

	sess.Register(caplets.NewCapletsModule(sess))
	sess.Register(update.NewUpdateModule(sess))
	sess.Register(ui.NewUIModule(sess))
}
