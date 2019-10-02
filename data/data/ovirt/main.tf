provider "ovirt" {
  url      = var.ovirt_url
  username = var.ovirt_username
  password = var.ovirt_password
}

module "bootstrap" {
  source = "./bootstrap"
  ovirt_cluster_id   = var.ovirt_cluster_id
  ovirt_template_id  = var.ovirt_template_id
  ocp_cluster_name   = var.ocp_cluster_name
  ignition_bootstrap = var.ignition_bootstrap
}

resource "ovirt_vm" "master0"  {
  name        = "ocp-cluster-${var.ocp_cluster_name}-master-0"
  cluster_id  = var.ovirt_cluster_id
  template_id = var.ovirt_template_id
  memory      = "8192"
  cores       = "4"

  initialization {
    host_name     = "master-0.local"
    custom_script = var.ignition_master
  }
}

resource "ovirt_vm" "master1"  {
  name        = "ocp-cluster-${var.ocp_cluster_name}-master-1"
  cluster_id  = var.ovirt_cluster_id
  template_id = var.ovirt_template_id
  memory      = "8192"
  cores       = "4"

  initialization {
    host_name     = "master-1.local"
    custom_script = var.ignition_master
  }
}

resource "ovirt_vm" "master2"  {
  name        = "ocp-cluster-${var.ocp_cluster_name}-master-2"
  cluster_id  = var.ovirt_cluster_id
  template_id = var.ovirt_template_id
  memory      = "8192"
  cores       = "4"

  initialization {
    host_name     = "master-2.local"
    custom_script = var.ignition_master
  }
}

