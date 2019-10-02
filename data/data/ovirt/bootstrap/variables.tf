variable "ignition_bootstrap" {
  description = "The ID of storage domain"
  default     = ""
}

variable "ovirt_cluster_id" {
  type = string
  default = ""
  description = "The name of cluster"
}

variable "ocp_cluster_name" {
  description = "The ID of Openshift cluster"
}

variable "ovirt_template_id" {
  type = string
  default = ""
  description = "The ID of cluster"
}
