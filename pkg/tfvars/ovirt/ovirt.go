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
	StorageDomainId   string `json:"ovirt_storage_domain_id,omitempty"`
	ClusterId         string `json:"ovirt_cluster_id,omitempty"`
	OcpClusterName    string `json:"ocp_cluster_name,omitempty"`
	TemplateId        string `json:"ovirt_template_id,omitempty"`
}

// TFVars generates ovirt-specific Terraform variables.
func TFVars(
	engineUrl string,
	engineUser string,
	enginePass string,
	engineCafile string,
	storageDomainId string,
	clusterId string,
	ocpClusterName string,
	templateId string) ([]byte, error) {

	cfg := config{
		Url:               engineUrl,
		Username:          engineUser,
		Password:          enginePass,
		Cafile:            engineCafile,
		StorageDomainId:   storageDomainId,
		ClusterId:         clusterId,
		OcpClusterName:    ocpClusterName,
		TemplateId:        templateId,
	}

	return json.MarshalIndent(cfg, "", "  ")
}
