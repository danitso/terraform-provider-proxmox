resource "proxmox_virtual_environment_role" "example" {
  privileges = [
    "VM.Monitor",
  ]
  role_id = "terraform-provider-proxmox-example"
}

output "resource_proxmox_virtual_environment_role_example" {
  value = proxmox_virtual_environment_role.example
}
