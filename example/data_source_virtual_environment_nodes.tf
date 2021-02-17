data "proxmox_virtual_environment_nodes" "example" {}

output "data_proxmox_virtual_environment_nodes_example" {
  value = data.proxmox_virtual_environment_nodes.example
}
