resource "proxmox_virtual_environment_pool" "example" {
  comment = "Managed by Terraform"
  pool_id = "terraform-provider-proxmox-example"
}

output "resource_proxmox_virtual_environment_pool_example" {
  value = proxmox_virtual_environment_pool.example
}
