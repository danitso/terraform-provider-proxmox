---
layout: home
title: Introduction
nav_order: 1
---

# Proxmox Provider

This provider for [Terraform](https://www.terraform.io/) is used for interacting with resources supported by [Proxmox](https://www.proxmox.com/en/). The provider needs to be configured with the proper endpoints and credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```
provider "proxmox" {
  virtual_environment {
    endpoint = "https://10.0.0.2"
    username = "root@pam"
    password = "the-password-set-during-installation-of-proxmox-ve"
    insecure = true
  }
}
```

```
provider "proxmox" {
  virtual_environment {
    endpoint = "https://10.0.0.2"
    insecure = true

    token = {
      id     = "root@pam!automation"
      secret = "a2d3b007-07a5-46d1-b52d-3ee0a15ffadd"
    }
  }
}
```

## Authentication

The Proxmox provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

* Static credentials
* Environment variables

### Static credentials

> Warning: Hard-coding credentials into any Terraform configuration is not recommended, and risks secret leakage should this file ever be committed to a public version control system.

Static credentials can be provided by adding the `username` and `password` arguments in-line in the Proxmox provider block:

```
provider "proxmox" {
  virtual_environment {
    username = "username@realm"
    password = "a-strong-password"
  }
}
```

A token can be provided the same way by adding an in-line `token` block instead:

```
provider "proxmox" {
  virtual_environment {
    token = {
      id     = "username@realm!my-token-name"
      secret = "4115ecd7-ba6f-41ef-bf42-2e19d0afaa88"
    }
  }
}
```

### Environment variables

You can provide your credentials via the `PROXMOX_VE_USERNAME` and `PROXMOX_VE_PASSWORD`, environment variables, representing your Proxmox username, realm and password, respectively:

```
provider "proxmox" {
  virtual_environment {}
}
```

Usage:

```sh
$ export PROXMOX_VE_USERNAME="username@realm"
$ export PROXMOX_VE_PASSWORD="a-strong-password"
$ terraform plan
```

## Argument Reference

In addition to [generic provider arguments](https://www.terraform.io/docs/configuration/providers.html) (e.g. `alias` and `version`), the following arguments are supported in the Proxmox `provider` block:

* `virtual_environment` - (Optional) The Proxmox Virtual Environment configuration.
    * `endpoint` - (Required) The endpoint for the Proxmox Virtual Environment API (can also be sourced from `PROXMOX_VE_ENDPOINT`).
    * `insecure` - (Optional) Whether to skip the TLS verification step (can also be sourced from `PROXMOX_VE_INSECURE`). If omitted, defaults to `false`.
    * `otp` - (Optional) The one-time password for the Proxmox Virtual Environment API (can also be sourced from `PROXMOX_VE_OTP`).
    * `password` - (Optional) The password for the Proxmox Virtual Environment API (can also be sourced from `PROXMOX_VE_PASSWORD`).
    * `token` - (Optional) The token for the Proxmox Virtual Environment API.
        * `id` - (Optional) The token identifier for the Proxmox Virtual Environment API (can also be sourced from `PROXMOX_VE_TOKEN_ID`).
        * `secret` - (Optional) The token secret for the Proxmox Virtual Environment API (can also be sourced from `PROXMOX_VE_TOKEN_SECRET`).
    * `username` - (Optional) The username and realm for the Proxmox Virtual Environment API (can also be sourced from `PROXMOX_VE_USERNAME`).
