package defaults

import (
	"github.com/openshift/installer/pkg/types/ovirt"
)

// SetPlatformDefaults sets the defaults for the platform.
func SetPlatformDefaults(p *ovirt.Platform) {
	if p.URI == "" {
		p.URI = DefaultURI
	}
	if p.Network == nil {
		p.Network = &ovirt.Network{}
	}
	SetNetworkDefaults(p.Network)
}
