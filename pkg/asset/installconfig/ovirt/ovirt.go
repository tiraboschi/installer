package ovirt

import (
	"github.com/openshift/installer/pkg/types/ovirt"
	"gopkg.in/AlecAivazis/survey.v1"
)

// Platform collects ovirt-specific configuration.
func Platform() (*ovirt.Platform, error) {
	var url string
	var username string
	var password string
	var apiVIP string

	err := survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Enter oVirt's api endpoint URL",
				Help:    "oVirt engine api url, for example: https://ovirt-engine-fqdn/ovirt-engine/api",
				Default: "https://rgolan.usersys.redhat.com:8443/ovirt-engine/api",
			},
			Validate: survey.ComposeValidators(survey.Required),
		},
	}, &url)
	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Enter the engine username",
				Help:    "The user must have permissions to create VMs, Disks, on the Storage Domain with the name matching the openshift cluster",
				Default: "admin@internal",
			},
			Validate: survey.ComposeValidators(survey.Required),
		},
	}, &username)
	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Enter password",
				Help:    "",
				Default: "",
			},
			Validate: survey.ComposeValidators(survey.Required),
		},
	}, &password)
	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Enter an API Virtual IP",
				Help:    "Make sure the ip is not used by any party",
				Default: "",
			},
			Validate: survey.ComposeValidators(survey.Required),
		},
	}, &apiVIP)

	if err != nil {
		return nil, err
	}

	return &ovirt.Platform{
		Url:      url,
		Username: username,
		Password: password,
		ApiVIP:   apiVIP,
		Cafile:   "",
	}, nil
}

