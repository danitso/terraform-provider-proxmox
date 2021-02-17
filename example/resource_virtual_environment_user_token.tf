resource "proxmox_virtual_environment_user_token" "example" {
  acl {
    path      = "/vms/${proxmox_virtual_environment_vm.example.id}"
    propagate = true
    role_id   = "PVEVMAdmin"
  }

  comment  = "Managed by Terraform"
  name     = "automation"
  user_id  = proxmox_virtual_environment_user.example.id
}

output "resource_proxmox_virtual_environment_user_token_example" {
  value = proxmox_virtual_environment_user_token.example
}
