data "proxmox_virtual_environment_hosts" "example" {
  node_name = data.proxmox_virtual_environment_nodes.example.names[0]
}

output "data_proxmox_virtual_environment_hosts_example" {
  value = data.proxmox_virtual_environment_hosts.example
}
