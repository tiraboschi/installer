package defaults

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/openshift/installer/pkg/types/ovirt"
)

func defaultPlatform() *ovirt.Platform {
	n := &ovirt.Network{}
	SetNetworkDefaults(n)
	return &ovirt.Platform{
		URI:     DefaultURI,
		Network: n,
	}
}

func TestSetPlatformDefaults(t *testing.T) {
	cases := []struct {
		name     string
		platform *ovirt.Platform
		expected *ovirt.Platform
	}{
		{
			name:     "empty",
			platform: &ovirt.Platform{},
			expected: defaultPlatform(),
		},
		{
			name: "URI present",
			platform: &ovirt.Platform{
				URI: "test-uri",
			},
			expected: func() *ovirt.Platform {
				p := defaultPlatform()
				p.URI = "test-uri"
				return p
			}(),
		},
		{
			name: "Network present",
			platform: &ovirt.Platform{
				Network: func() *ovirt.Network {
					n := &ovirt.Network{}
					SetNetworkDefaults(n)
					n.IfName = "test-if"
					return n
				}(),
			},
			expected: func() *ovirt.Platform {
				p := defaultPlatform()
				p.Network = &ovirt.Network{}
				SetNetworkDefaults(p.Network)
				p.Network.IfName = "test-if"
				return p
			}(),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			SetPlatformDefaults(tc.platform)
			assert.Equal(t, tc.expected, tc.platform, "unexpected platform")
		})
	}
}
