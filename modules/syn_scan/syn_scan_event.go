package syn_scan

import (
	"github.com/bdsoftpro/wifi/network"
	"github.com/bdsoftpro/wifi/session"
)

type SynScanEvent struct {
	Address string
	Host    *network.Endpoint
	Port    int
}

func NewSynScanEvent(address string, h *network.Endpoint, port int) SynScanEvent {
	return SynScanEvent{
		Address: address,
		Host:    h,
		Port:    port,
	}
}

func (e SynScanEvent) Push() {
	session.I.Events.Add("syn.scan", e)
	session.I.Refresh()
}
