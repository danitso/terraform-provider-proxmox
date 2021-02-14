data "proxmox_virtual_environment_cluster_aliases" "example" {
  depends_on = [proxmox_virtual_environment_cluster_alias.example]
}

output "data_proxmox_virtual_environment_cluster_aliases_example" {
  value = data.proxmox_virtual_environment_cluster_aliases.example
}
