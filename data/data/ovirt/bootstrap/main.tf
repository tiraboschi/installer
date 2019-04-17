resource "ovirt_vm" "bootstrap" {
  name        = "${var.cluster_name}-bootstrap"
  memory      = "8192"
  cores       = "4"
  cluster_id  = var.cluster_id
  template_id = var.template_id

  initialization {
    custom_script = var.ignition_bootstrap
  }
}
