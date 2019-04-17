package ovirt

import (
	"github.com/openshift/installer/pkg/types/ovirt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_validateAuth(t *testing.T) {
	tests := []struct {
		url           string
		username      string
		password      string
		insecure      bool
		cafile        string
		expectSuccess bool
	}{{
		url:           "https://rgolan.usersys.redhat.com:8443/ovirt-engine/api",
		username:      "admin@internal",
		password:      "123",
		insecure:      false,
		cafile:        "",
		expectSuccess: true,
	},
		{
			url:           "https://nonexisting",
			username:      "foo",
			password:      "bar",
			insecure:      false,
			cafile:        "",
			expectSuccess: false,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			p := ovirt.Platform{
				Url:      test.url,
				Username: test.username,
				Password: test.password,
				Cafile:   test.cafile,
				Insecure: test.insecure,
			}

			validationFunc := Authenticated(&p)
			got := validationFunc(p.Password)
			assert.Equal(t, test.expectSuccess, got == nil, "got this %s", got )
			t.Log(got)
		})
	}
}
