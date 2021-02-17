data "proxmox_virtual_environment_pool" "example" {
  pool_id = proxmox_virtual_environment_pool.example.id
}

output "data_proxmox_virtual_environment_pool_example" {
  value = data.proxmox_virtual_environment_pool.example
}
