resource "proxmox_virtual_environment_group" "example" {
  acl {
    path    = "/vms/${proxmox_virtual_environment_vm.example.id}"
    role_id = proxmox_virtual_environment_role.example.id
  }

  comment  = "Managed by Terraform"
  group_id = "terraform-provider-proxmox-example"
}

output "resource_proxmox_virtual_environment_group_example" {
  value = proxmox_virtual_environment_group.example
}
