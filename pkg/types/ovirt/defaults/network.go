package defaults

import (
	"github.com/openshift/installer/pkg/ipnet"

	"github.com/openshift/installer/pkg/types/ovirt"
)

const (
	defaultIfName = "tt0"
)

var (
	// DefaultMachineCIDR is the ovirt default IP address space from
	// which to assign machine IPs.
	DefaultMachineCIDR = ipnet.MustParseCIDR("192.168.126.0/24")
)

// SetNetworkDefaults sets the defaults for the network.
func SetNetworkDefaults(n *ovirt.Network) {
	if n.IfName == "" {
		n.IfName = defaultIfName
	}
}
