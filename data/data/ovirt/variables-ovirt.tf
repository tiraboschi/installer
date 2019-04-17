variable "bootstrap_dns" {
  type = string
  default     = true
  description = "Whether to include DNS entries for the bootstrap node or not."
}

variable "ovirt_url" {
  type = string
  default = ""
  description = "The oVirt engine URL"
}

variable "ovirt_username" {
  type = string
  default = ""
  description = "The name of user to access oVirt engine API"
}

variable "ovirt_password" {
  type = string
  default = ""
  description = "The plain password of user to access oVirt engine API"
}

variable "cluster_name" {
  type = string
  default = "blue"
  description = "The name of cluster"
}

//variable "cluster_id" {
//  description = "The ID of cluster"
//}

variable "storage_domain_id" {
  type = string
  default = ""
  description = "The ID of storage domain"
}

variable "storage_domain_name" {
  type = string
  description = "The name of storage domain"
  default     = "nfs"
}

variable "template_name" {
  type = string
  default     = "rhcos"
  description = "The ID of VM template"
}

variable "template_id" {
  type = string
  default = ""
  description = "The ID of cluster"
}

