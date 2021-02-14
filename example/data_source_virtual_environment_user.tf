data "proxmox_virtual_environment_user" "example" {
  user_id = proxmox_virtual_environment_user.example.id
}

output "data_proxmox_virtual_environment_user_example" {
  value = data.proxmox_virtual_environment_user.example
}
