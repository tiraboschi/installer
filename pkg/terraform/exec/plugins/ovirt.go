package plugins

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/ovirt/terraform-provider-ovirt/ovirt"
)

func init() {
	exec := func() {
	
		plugin.Serve(&plugin.ServeOpts{
			ProviderFunc: ovirt.Provider,
		})
	}
	KnownPlugins["terraform-provider-ovirt"] = exec
}
