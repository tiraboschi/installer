package types

import (
	"github.com/openshift/installer/pkg/types/ovirt"
	"sort"
)

func init() {
	PlatformNames = append(PlatformNames, ovirt.Name)
	sort.Strings(PlatformNames)
}
