---
layout: "appgate"
page_title: "Provider: AppGate"
description: |-

---

# AppGate

## Example Usage

```hcl
# Configure the AppGate Provider
provider "appgate" {
  username = "admin"
  password = "admin"
  url      = "https://controller.devops:444/admin"
  provider = "local"
  insecure = true
}


```

## Authentication

The AppGate provider offers a flexible means of providing credentials for
authentication. The following methods are supported, in this order, and
explained below:

- Static credentials
- Environment variables

### Static credentials

!> **Warning:** Hard-coding credentials into any Terraform configuration is not
recommended, and risks secret leakage should this file ever be committed to a
public version control system.

Static credentials can be provided by adding an `username` and `password`
in-line in the AppGate provider block:

Usage:

```hcl
provider "appgate" {
  username = "admin"
  password = "admin"
  provider = "local"
}
```

### Environment variables


```hcl
provider "appgate" {}
```

Usage:

```sh
$ export APPGATE_ADDRESS="https://controller.devops:444/admin"
$ export APPGATE_USERNAME="admin"
$ export APPGATE_PASSWORD="admin"
$ export APPGATE_PROVIDER="local"
$ export APPGATE_INSECURE="true"
$ terraform plan
```


## Argument Reference

In addition to [generic `provider` arguments](https://www.terraform.io/docs/configuration/providers.html)
(e.g. `alias` and `version`), the following arguments are supported in the AppGate
 `provider` block:

* `username` - (Optional) This is the AppGate username. It must be provided, but
  it can also be sourced from the `APPGATE_USERNAME` environment variable.

* `password` - (Optional) This is the AppGate password. It must be provided, but
  it can also be sourced from the `APPGATE_PASSWORD` environment variable.

* `provider` - (Optional) This is the AppGate provider. It must be provided, but
  it can also be sourced from the `APPGATE_PROVIDER` environment variables.

* `insecure` - (Optional) Whether server should be accessed without verifying the TLS certificate. As the name suggests this is insecure and should not be used beyond experiments, accessing local (non-production) GHE instance etc. There is a number of ways to obtain trusted certificate for free, e.g. from Let's Encrypt. Such trusted certificate does not require this option to be enabled. Defaults to `true`.