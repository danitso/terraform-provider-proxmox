data "proxmox_virtual_environment_users" "example" {
  depends_on = [proxmox_virtual_environment_user.example]
}

output "data_proxmox_virtual_environment_users_example" {
  value = data.proxmox_virtual_environment_users.example
}
