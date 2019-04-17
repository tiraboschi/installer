package ovirt

// Metadata contains ovirt metadata (e.g. for uninstalling the cluster).
type Metadata struct {
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	Cafile   string `json:"cafile"`
	ApiVIP   string `json:"api_vip"`
	DnsVIP   string `json:"dns_vip"`
}
