package ovirt

import (
	"github.com/openshift/installer/pkg/destroy"
)

func init() {
	destroy.Registry["ovirt"] = New
}
