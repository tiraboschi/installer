package defaults

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/openshift/installer/pkg/types/ovirt"
)

func defaultNetwork() *ovirt.Network {
	return &ovirt.Network{
		IfName: defaultIfName,
	}
}

func TestSetNetworkDefaults(t *testing.T) {
	cases := []struct {
		name     string
		network  *ovirt.Network
		expected *ovirt.Network
	}{
		{
			name:     "empty",
			network:  &ovirt.Network{},
			expected: defaultNetwork(),
		},
		{
			name: "IfName present",
			network: &ovirt.Network{
				IfName: "test-if",
			},
			expected: func() *ovirt.Network {
				n := defaultNetwork()
				n.IfName = "test-if"
				return n
			}(),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			SetNetworkDefaults(tc.network)
			assert.Equal(t, tc.expected, tc.network, "unexpected network")
		})
	}
}
