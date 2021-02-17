resource "proxmox_virtual_environment_cluster_alias" "example" {
  name    = "example"
  cidr    = "192.168.0.0/23"
  comment = "Managed by Terraform"
}

output "proxmox_virtual_environment_cluster_alias_example" {
  value = proxmox_virtual_environment_cluster_alias.example
}
