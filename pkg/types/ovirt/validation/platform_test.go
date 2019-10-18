package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/openshift/installer/pkg/types/ovirt"
)

func validPlatform() *ovirt.Platform {
	return &ovirt.Platform{
		URI: "qemu+tcp://192.168.122.1/system",
		Network: &ovirt.Network{
			IfName: "tt0",
		},
	}
}

func TestValidatePlatform(t *testing.T) {
	cases := []struct {
		name     string
		platform *ovirt.Platform
		valid    bool
	}{
		{
			name:     "minimal",
			platform: validPlatform(),
			valid:    true,
		},
		{
			name: "invalid uri",
			platform: func() *ovirt.Platform {
				p := validPlatform()
				p.URI = "bad-uri"
				return p
			}(),
			valid: false,
		},
		{
			name: "missing network",
			platform: func() *ovirt.Platform {
				p := validPlatform()
				p.Network = nil
				return p
			}(),
			valid: false,
		},
		{
			name: "missing interface name",
			platform: func() *ovirt.Platform {
				p := validPlatform()
				p.Network.IfName = ""
				return p
			}(),
			valid: false,
		},
		{
			name: "valid machine pool",
			platform: func() *ovirt.Platform {
				p := validPlatform()
				p.DefaultMachinePlatform = &ovirt.MachinePool{}
				return p
			}(),
			valid: true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidatePlatform(tc.platform, field.NewPath("test-path")).ToAggregate()
			if tc.valid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
