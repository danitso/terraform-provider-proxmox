---
layout: page
title: proxmox_virtual_environment_user_token
permalink: /resources/virtual_environment_user_token
nav_order: 12
parent: Resources
subcategory: Virtual Environment
---

# Resource: proxmox_virtual_environment_user_token

Manages a user token.

## Example Usage

```
resource "proxmox_virtual_environment_user" "operations_automation" {
  acl {
    path      = "/vms/1234"
    propagate = true
    role_id   = "${proxmox_virtual_environment_role.operations_monitoring.role_id}"
  }

  comment  = "Managed by Terraform"
  password = "a-strong-password"
  user_id  = "operations-automation@pve"
}

resource "proxmox_virtual_environment_role" "operations_monitoring" {
  role_id = "operations-monitoring"

  privileges = [
    "VM.Monitor",
  ]
}
```

## Argument Reference

* `acl` - (Optional) The access control list (multiple blocks supported).
    * `path` - The path.
    * `propagate` - Whether to propagate to child paths.
    * `role_id` - The role identifier.
* `comment` - (Optional) The token comment.
* `expiration_date` - (Optional) The token's expiration date (RFC 3339).
* `name` - (Optional) The token's name.
* `privilege_separation` - (Optional) Whether the privileges for the token differs from the account privileges.
* `user_id` - (Required) The user identifier.

## Attribute Reference

* `secret` - The token secret.
