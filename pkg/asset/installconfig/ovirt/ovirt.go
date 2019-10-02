package ovirt

import (
	"fmt"
	"github.com/openshift/installer/pkg/types/ovirt"
	ovirtsdk4 "github.com/ovirt/go-ovirt"
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
			Validate: survey.ComposeValidators(survey.Required, Authenticated(&p)),
		},
	}, &p.Password)

	c, err := ovirtsdk4.NewConnectionBuilder().
		URL(p.Url).
		Username(p.Username).
		Password(p.Password).
		CAFile(p.Cafile).
		Insecure(false).
		Build()

	if err != nil {
		return nil, err
	}
	defer c.Close()
	err = c.Test()
	if err != nil {
		return nil, err
	}

	var clusterName string
	var clusterByNames map[string]*ovirtsdk4.Cluster = make(map[string]*ovirtsdk4.Cluster)
    var clusterNames []string
	systemService := c.SystemService()
	response, err := systemService.ClustersService().List().Send()
	if err != nil {
		return nil, err
	}
	clusters, ok := response.Clusters()
	if !ok {
		return nil, fmt.Errorf("there are no available clusters")
	}

	for _, cluster := range clusters.Slice() {
		clusterByNames[cluster.MustName()]= cluster
		clusterNames = append(clusterNames, cluster.MustName())
	}
	err = survey.AskOne(&survey.Select{
		Message: "Pick the oVirt cluster",
		Help:    "The oVirt cluster under which the VMs will be created.",
		Options: clusterNames,
	},
		&clusterName,
		func(ans interface{}) error {
			choice := ans.(string)
			sort.Strings(clusterNames)
			i := sort.SearchStrings(clusterNames, choice)
			if i == len(clusterNames) || clusterNames[i] != choice {
				return fmt.Errorf("invalid cluster %s", choice)
			}
			cl, ok := clusterByNames[choice]
			if !ok {
				return fmt.Errorf("cannot find a cluster id for the cluster name %s", clusterName)
			}
			p.ClusterId = cl.MustId()
			return nil
		})

	var templateName string
	var templateByNames map[string]*ovirtsdk4.Template = make(map[string]*ovirtsdk4.Template)
	var templateNames []string

	templateList, err := systemService.TemplatesService().List().Search("cluster=" + clusterName).Send()
	if err != nil {
		return nil, err
	}
	templates, ok := templateList.Templates()
	if !ok {
		return nil, fmt.Errorf("could not fetch the VM templates of cluster %s", clusterName)
	}
	for _, tmpl := range templates.Slice() {
		templateByNames[tmpl.MustName()] = tmpl
		templateNames = append(templateNames, tmpl.MustName())
	}

	err = survey.AskOne(&survey.Select{
		Message: "Pick a VM template",
		Help: "The VM template will be used to create every node in the cluster",
		Options: templateNames,
	},
		&templateName,
		func(ans interface{}) error {
			choice := ans.(string)
			sort.Strings(templateNames)
			i := sort.SearchStrings(templateNames, choice)
			if i == len(templateNames) || templateNames[i] != choice {
				return fmt.Errorf("invalid template %s", choice)
			}
			tmpl, _ := templateByNames[choice]
			p.TemplateId = tmpl.MustId()
			return nil
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
				Message: "Enter the internal DNS Virtual IP",
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
