data "proxmox_virtual_environment_roles" "example" {
  depends_on = [proxmox_virtual_environment_role.example]
}

output "data_proxmox_virtual_environment_roles_example" {
  value = data.proxmox_virtual_environment_roles.example
}
