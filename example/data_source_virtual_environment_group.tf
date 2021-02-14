data "proxmox_virtual_environment_group" "example" {
  group_id = proxmox_virtual_environment_group.example.id
}

output "data_proxmox_virtual_environment_group_example" {
  value = data.proxmox_virtual_environment_group.example
}
