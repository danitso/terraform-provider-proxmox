resource "proxmox_virtual_environment_time" "example" {
  node_name = data.proxmox_virtual_environment_time.example.node_name
  time_zone = data.proxmox_virtual_environment_time.example.time_zone
}

output "resource_proxmox_virtual_environment_time_example" {
  value = proxmox_virtual_environment_time.example
}
