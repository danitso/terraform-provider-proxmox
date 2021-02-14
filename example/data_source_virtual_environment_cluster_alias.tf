data "proxmox_virtual_environment_cluster_alias" "example" {
  depends_on = [proxmox_virtual_environment_cluster_alias.example]

  name = proxmox_virtual_environment_cluster_alias.example.name
}

output "data_proxmox_virtual_environment_cluster_alias_example" {
  value = proxmox_virtual_environment_cluster_alias.example
}
