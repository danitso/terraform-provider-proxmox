#===============================================================================
# Cloud Config (cloud-init)
#===============================================================================

resource "proxmox_virtual_environment_file" "cloud_config" {
  content_type = "snippets"
  datastore_id = element(data.proxmox_virtual_environment_datastores.example.datastore_ids, index(data.proxmox_virtual_environment_datastores.example.datastore_ids, "local"))
  node_name    = data.proxmox_virtual_environment_datastores.example.node_name

  source_raw {
    data = <<EOF
#cloud-config
chpasswd:
  list: |
    ubuntu:example
  expire: false
hostname: terraform-provider-proxmox-example
packages:
  - qemu-guest-agent
users:
  - default
  - name: ubuntu
    groups: sudo
    shell: /bin/bash
    ssh-authorized-keys:
      - ${trimspace(tls_private_key.example.public_key_openssh)}
    sudo: ALL=(ALL) NOPASSWD:ALL
    EOF

    file_name = "terraform-provider-proxmox-example-cloud-config.yaml"
  }
}

output "resource_proxmox_virtual_environment_file_cloud_config" {
  value = proxmox_virtual_environment_file.cloud_config
}

#===============================================================================
# Ubuntu Cloud Image
#===============================================================================

resource "proxmox_virtual_environment_file" "ubuntu_cloud_image" {
  content_type = "iso"
  datastore_id = element(data.proxmox_virtual_environment_datastores.example.datastore_ids, index(data.proxmox_virtual_environment_datastores.example.datastore_ids, "local"))
  node_name    = data.proxmox_virtual_environment_datastores.example.node_name

  source_file {
    path = "https://cloud-images.ubuntu.com/bionic/current/bionic-server-cloudimg-amd64.img"
  }
}

output "resource_proxmox_virtual_environment_file_ubuntu_cloud_image" {
  value = proxmox_virtual_environment_file.ubuntu_cloud_image
}

#===============================================================================
# Ubuntu Container Template
#===============================================================================

resource "proxmox_virtual_environment_file" "ubuntu_container_template" {
  content_type = "vztmpl"
  datastore_id = element(data.proxmox_virtual_environment_datastores.example.datastore_ids, index(data.proxmox_virtual_environment_datastores.example.datastore_ids, "local"))
  node_name    = data.proxmox_virtual_environment_datastores.example.node_name

  source_file {
    path = "http://download.proxmox.com/images/system/ubuntu-18.04-standard_18.04.1-1_amd64.tar.gz"
  }
}

output "resource_proxmox_virtual_environment_file_ubuntu_container_template" {
  value = proxmox_virtual_environment_file.ubuntu_container_template
}
