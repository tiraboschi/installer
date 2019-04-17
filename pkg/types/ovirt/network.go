package ovirt

// Network is the configuration of the ovirt network.
type Network struct {
	// +optional
	// Default is tt0.
	IfName string `json:"if,omitempty"`
}
