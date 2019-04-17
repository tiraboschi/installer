// Package ovirt extracts ovirt metadata from install configurations.
package ovirt

import (
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/ovirt"
)

// Metadata converts an install configuration to ovirt metadata.
func Metadata(config *types.InstallConfig) *ovirt.Metadata {
	return &ovirt.Metadata{
		Url:        config.Platform.Ovirt.Url,
		Username:   config.Platform.Ovirt.Username,
		Password:   config.Platform.Ovirt.Password,
		Cafile:     config.Platform.Ovirt.Cafile,
		ApiVIP:     config.Platform.Ovirt.ApiVIP,
		DnsVIP:     config.Platform.Ovirt.DnsVIP,
		IngressVIP: config.Platform.Ovirt.IngressVIP,
	}
}
