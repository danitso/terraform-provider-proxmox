data "proxmox_virtual_environment_role" "example" {
  role_id = proxmox_virtual_environment_role.example.id
}

output "data_proxmox_virtual_environment_role_example" {
  value = data.proxmox_virtual_environment_role.example
}
