// Package ovirt contains ovirt-specific Terraform-variable logic.
package ovirt

import (
	"encoding/json"
)

type config struct {
	Url               string `json:"ovirt_url,omitempty"`
	Username          string `json:"ovirt_username,omitempty"`
	Password          string `json:"ovirt_password,omitempty"`
	Cafile            string `json:"ovirt_cafile,omitempty"`
	StorageDomainId   string `json:"storage_domain_id,omitempty"`
	StorageDomainName string `json:"storage_domain_name,omitempty"`
	ClusterId         string `json:"cluster_id,omitempty"`
	ClusterName       string `json:"cluster_name,omitempty"`
	TemplateId        string `json:"template_id,omitempty"`
	TemplateName      string `json:"template_name,omitempty"`
	Image             string `json:"os_image,omitempty"`
}

// TFVars generates ovirt-specific Terraform variables.
func TFVars(
	engineUrl string,
	engineUser string,
	enginePass string,
	engineCafile string,
	storageDomainName string,
	clusterName string,
	templateName string,
	osImage string,
	masterCount int) ([]byte, error) {

	cfg := config{
		Url:               engineUrl,
		Username:          engineUser,
		Password:          enginePass,
		Cafile:            engineCafile,
		StorageDomainId:   "d787bf6b-fae1-4a3e-b773-2ac466599d29",
		StorageDomainName: storageDomainName,
		ClusterId:         "5c8f6906-f14b-43ee-83df-5f800f36eb70",
		ClusterName:       clusterName,
		TemplateId:        "31e0997c-e4a8-45a2-9f0f-8b689056be37",
		TemplateName:      templateName,
		//IfName:      bridge,
		//BootstrapIP: bootstrapIP.String(),
		//MasterIPs:   masterIPs,
	}

	return json.MarshalIndent(cfg, "", "  ")
}
