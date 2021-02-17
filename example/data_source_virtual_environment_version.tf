data "proxmox_virtual_environment_version" "example" {}

output "data_proxmox_virtual_environment_version_example" {
  value = data.proxmox_virtual_environment_version.example
}
