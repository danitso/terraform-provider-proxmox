---
layout: home
title: Introduction
nav_order: 1
---

# Proxmox Provider

This provider for [Terraform](https://www.terraform.io/) is used to interact with resources supported by [Proxmox](https://www.proxmox.com/en/). The provider needs to be configured with the proper endpoints and credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```
provider "proxmox" {
  virtual_environment {
    endpoint = "https://10.0.0.2:8006"
    username = "root@pam"
    password = "the-password-set-during-installation-of-proxmox-ve"
    insecure = true
  }
}
```

## Installation

You can install the latest release of the provider using either Git Bash or regular Bash:

```sh
$ export PROVIDER_PLATFORM="$([[ "$OSTYPE" =~ ^msys|cygwin$ ]] && echo "windows" || ([[ "$OSTYPE" == "darwin"* ]] && echo "darwin" || ([[ "$OSTYPE" == "linux"* ]] && echo "linux" || echo "unsupported")))"
$ export PROVIDER_VERSION="$(curl -L -s -H 'Accept: application/json' https://github.com/danitso/terraform-provider-proxmox/releases/latest | sed -e 's/.*"tag_name":"\([^"]*\)".*/\1/')"
$ export PLUGINS_PATH="$([[ "$PROVIDER_PLATFORM" == "windows" ]] && cygpath -u "$APPDATA" || echo "$HOME")/terraform.d/plugins"
$ mkdir -p "$PLUGINS_PATH"
$ curl -o "${PLUGINS_PATH}/terraform-provider-proxmox_v${PROVIDER_VERSION}.zip" -sL "https://github.com/danitso/terraform-provider-proxmox/releases/download/${PROVIDER_VERSION}/terraform-provider-proxmox_v${PROVIDER_VERSION}-custom_${PROVIDER_PLATFORM}_amd64.zip"
$ unzip -o -d "$PLUGINS_PATH" "${PLUGINS_PATH}/terraform-provider-proxmox_v${PROVIDER_VERSION}.zip"
$ rm "${PLUGINS_PATH}/terraform-provider-proxmox_v${PROVIDER_VERSION}.zip"
```

You can also install it manually by following the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-plugins). You can download the latest release from the [releases](https://github.com/danitso/terraform-provider-proxmox/releases) page.

## Authentication

The Proxmox provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

* Static credentials
* Environment variables

### Static credentials

Warning: Hard-coding credentials into any Terraform configuration is not recommended, and risks secret leakage should this file ever be committed to a public version control system.
{: .label .label-red }

Static credentials can be provided by adding a `username` and `password` in-line in the Proxmox provider block:

```
provider "proxmox" {
  virtual_environment {
    username = "username@realm"
    password = "a-strong-password"
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

## Arguments Reference

In addition to [generic provider arguments](https://www.terraform.io/docs/configuration/providers.html) (e.g. `alias` and `version`), the following arguments are supported in the Proxmox `provider` block:

* `virtual_environment` - (Optional) The Proxmox Virtual Environment configuration.
    * `endpoint` - (Required) The endpoint for the Proxmox Virtual Environment API (can also be sourced from `PROXMOX_VE_ENDPOINT`). Note that the `/api2/json` part will be automatically added by the provider so there is no need to specify it.
    * `insecure` - (Optional) Whether to skip the TLS verification step (can also be sourced from `PROXMOX_VE_INSECURE`). If omitted, defaults to `false`.
    * `otp` - (Optional) The one-time password for the Proxmox Virtual Environment API (can also be sourced from `PROXMOX_VE_OTP`).
    * `password` - (Required) The password for the Proxmox Virtual Environment API (can also be sourced from `PROXMOX_VE_PASSWORD`).
    * `username` - (Required) The username and realm for the Proxmox Virtual Environment API (can also be sourced from `PROXMOX_VE_USERNAME`).
