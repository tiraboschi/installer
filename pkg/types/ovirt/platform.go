package ovirt

// Platform stores all the global configuration that all
// machinesets use.
type Platform struct {
	// +optional
	Url               string `json:"ovirt_url,omitempty"`
	Username          string `json:"ovirt_username,omitempty"`
	Password          string `json:"ovirt_password,omitempty"`
	Cafile            string `json:"ovirt_cafile,omitempty"`
	Insecure		  bool   `json:"ovirt_insecure,omitempty"`
	StorageDomainName string `json:"storage_domain_name,omitempty"`
	ClusterId         string `json:"cluster_id,omitempty"`
	TemplateId        string `json:"template_id,omitempty"`

	// DefaultMachinePlatform is the default configuration used when
	// installing on ovirt for machine pools which do not define their
	// own platform configuration.
	// +optional
	// Default will set the image field to the latest RHCOS image.
	DefaultMachinePlatform *MachinePool `json:"defaultMachinePlatform,omitempty"`

	// Network
	// +optional
	Network         *Network `json:"network,omitempty"`

	// ApiVIP is an IP which will be served by bootstrap and then pivoted masters, using keepalived
	ApiVIP string `json:"api_vip,omitempty"`

	// DnsVIP is the IP of the internal DNS which will be operated by the cluster
	DnsVIP string `json:"dns_vip,omitempty"`

}
