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
	ClusterId         string `json:"ovirt_cluster_id,omitempty"`
	TemplateId        string `json:"ovirt_template_id,omitempty"`
}

// TFVars generates ovirt-specific Terraform variables.
func TFVars(
	engineUrl string,
	engineUser string,
	enginePass string,
	engineCafile string,
	clusterId string,
	templateId string) ([]byte, error) {

	cfg := config{
		Url:               engineUrl,
		Username:          engineUser,
		Password:          enginePass,
		Cafile:            engineCafile,
		ClusterId:         clusterId,
		TemplateId:        templateId,
	}

	return json.MarshalIndent(cfg, "", "  ")
}
