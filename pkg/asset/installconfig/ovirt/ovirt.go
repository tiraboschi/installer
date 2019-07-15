package ovirt

import (
	"fmt"
	"github.com/openshift/installer/pkg/types/ovirt"
	"gopkg.in/AlecAivazis/survey.v1"
	"sort"
)

// Platform collects ovirt-specific configuration.
func Platform() (*ovirt.Platform, error) {
	p := ovirt.Platform{}

	err := survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Enter oVirt's api endpoint URL",
				Help:    "oVirt engine api url, for example: https://ovirt-engine-fqdn/ovirt-engine/api",
				Default: "https://rgolan.usersys.redhat.com:8443/ovirt-engine/api",
			},
			Validate: survey.ComposeValidators(survey.Required),
		},
	}, &p.Url)

	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Enter the engine username",
				Help:    "The user must have permissions to create VMs, Disks, on the Storage Domain with the name matching the openshift cluster",
				Default: "admin@internal",
			},
			Validate: survey.ComposeValidators(survey.Required),
		},
	}, &p.Username)

	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Enter password",
				Help:    "",
				Default: "",
			},
			Validate: survey.ComposeValidators(survey.Required, Authenticated(p)),
		},
	}, &p.Password)

	c, err := ovirtsdk.NewConnectionBuilder().
		URL(p.Url).
		Username(p.Username).
		Password(p.Password).
		CAFile(p.Cafile).
		Insecure(p.Insecure).
		Build()

	if err != nil {
		return nil, err
	}
	err = c.Test()
	if err != nil {
		return nil, err
	}

	var clusterName string
	var clusterNames []string

	response, err := c.SystemService().ClustersService().List().Send()
	if err != nil {
		return nil, err
	}
	clusters, ok := response.Clusters()
	if !ok {
		return nil, fmt.Errorf("there are no available cluster under oVirt setup")
	}

	for _, cluster := range clusters.Slice() {
		clusterNames = append(clusterNames, cluster.MustName())
	}
	err = survey.AskOne(&survey.Select{
		Message: "oVirt cluster",
		Help:    "The oVirt cluster under which the VMs will be created.",
		Options: clusterNames,
	},
	&clusterName,
	func(ans interface{}) error {
		choice := ans.(string)
		i := sort.SearchStrings(clusterNames, choice)
		if i == len(clusterNames) || clusterNames[i] != choice {
			return fmt.Errorf("invalid cluster %s", choice)
		}
		for _, cluster := range clusters.Slice() {
			if cluster.MustName() == clusterName {
				p.ClusterId = cluster.MustId()
				return nil
			}
		}
		return fmt.Errorf("cannot find a cluster id for the cluster name %s", clusterName)
	})

	var templateName string
	var templateNames []string

	templateList, err := c.SystemService().TemplatesService().List().Search("cluster=" + clusterName).Send()
	if err != nil {
		return nil, err
	}
	templates := templateList.MustTemplates()

	err = survey.AskOne(
		&survey.Select{

		},
		&templateName,
		func (ans interface{}) error {
			choice := ans.(string)
			i := sort.SearchStrings(templateNames, choice)
			if i == len(templateNames) || templateNames[i] != choice {
				return fmt.Errorf("invalid template %s", choice)
			}
			for _, template := range templates.Slice() {
				if template.MustName() == templateName {
					p.TemplateId = template.MustId()
					return nil
				}
			}
			return fmt.Errorf("could not locate template %s", templateName)
		})


	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Enter the internal API Virtual IP",
				Help:    "Make sure the ip is not used by any party",
				Default: "",
			},
			Validate: survey.ComposeValidators(survey.Required),
		},
	}, &p.ApiVIP)
	err = survey.Ask([]*survey.Question{
		{
			Prompt: &survey.Input{
				Message: "Enter the internal API Virtual IP",
				Help:    "Make sure the ip is not used by any party",
				Default: "",
			},
			Validate: survey.ComposeValidators(survey.Required),
		},
	}, &p.DnsVIP)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

