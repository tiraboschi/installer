provider "ovirt" {
  url      = var.ovirt_url
  username = var.ovirt_username
  password = var.ovirt_password
}

module "bootstrap" {
  source             = "./bootstrap"
  storage_domain_id  = var.storage_domain_id
  cluster_id         = var.cluster_id
  template_id        = var.template_id
  ignition_bootstrap = var.ignition_bootstrap
}

resource "ovirt_vm" "master0"  {
  name        = "${var.cluster_name}-master-0"
  cluster_id  = var.cluster_id
  template_id = var.template_id
  memory      = "8192"
  cores       = "4"

  initialization {
    host_name     = "master-0.local"
    custom_script = var.ignition_master
  }
}

resource "ovirt_vm" "master1"  {
  name        = "${var.cluster_name}-master-1"
  cluster_id  = var.cluster_id
  template_id = var.template_id
  memory      = "8192"
  cores       = "4"

  initialization {
    host_name     = "master-1.local"
    custom_script = var.ignition_master
  }
}

resource "ovirt_vm" "master2"  {
  name        = "${var.cluster_name}-master-2"
  cluster_id  = var.cluster_id
  template_id = var.template_id
  memory      = "8192"
  cores       = "4"

  initialization {
    host_name     = "master-2.local"
    custom_script = var.ignition_master
  }
}

