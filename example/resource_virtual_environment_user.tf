resource "proxmox_virtual_environment_user" "example" {
  acl {
    path      = "/vms/${proxmox_virtual_environment_vm.example.id}"
    propagate = true
    role_id   = "PVEVMAdmin"
  }

  comment  = "Managed by Terraform"
  password = "Test1234!"
  user_id  = "terraform-provider-proxmox-example@pve"
}

output "resource_proxmox_virtual_environment_user_example" {
  value = proxmox_virtual_environment_user.example
}
