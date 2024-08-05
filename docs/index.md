---
organization: Akeyless Security
category: ["security"]
icon_url: "/images/plugins/akeyless-community/akeyless.svg"
brand_color: "#01D9C1"
display_name: "Akeyless"
short_name: "akeyless"
description: "Steampipe plugin to query items, auth methods, roles, etc from Akeyless."
og_description: Query Akeyless with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/akeyless-community/akeyless-social-graphic.png"
---

# Akeyless + Steampipe

[Akeyless](https://www.akeyless.io/) The Akeyless Platform is a unified secrets management system that enables you to store, protect, rotate, and dynamically create and manage credentials, certificates, and encryption keys.

[Steampipe](https://steampipe.io/) is an open source CLI for querying cloud APIs using SQL from [Turbot](https://turbot.com/)

List `roles` in your Akeyless account:

```sql
select
  role_name
from
  akeyless_role;
```

## Documentation

- **[Table definitions & examples →](https://hub.steampipe.io/plugins/akeyless/akeyless/tables)**

## Get started

### Install

Download and install the latest Akeyless plugin:

```shell
steampipe plugin install akeyless/akeyless
```

### Configuration

Installing the latest Akeyless plugin will create a config file (`~/.steampipe/config/akeyless.spc`) with a single connection named `akeyless`:

```hcl
connection "akeyless" {
  plugin = "akeyless"  

  # Defines the type of access. Supported types: "api_key", "aws_iam", "azure_ad", "gcp", "universal_identity", "k8s", "jwt".
  #access_type = ""

  # The access ID for authentication, required for all access types.
  #access_id = ""

  # The access key or secret, paired with access_id.
  #access_key = ""

  # API URL for the Akeyless Gateway, the default URL is https://api.akeyless.io.
  #api_url = ""

  # JSON Web Token for JWT-based authentication.
  #jwt = ""

  # User identity token for Universal Identity™ authentication.
  #uid_token = ""

  # Audience for GCP authentication.
  #gcp_audience = ""

  # Object ID for Azure AD authentication.
  #azure_object_id = ""

  # Kubernetes service account token for Kubernetes-based authentication.
  #k8s_service_account_token = ""

  # Name of the Kubernetes auth config.
  #k8s_auth_config_name = ""

  # CA certificate for TLS verification of the Akeyless Gateway.
  #gateway_ca_cert = ""
}
```

#### Authentication

Akeyless supports multiple authentication methods: `api_key`,`aws_iam`,`azure_ad`,`gcp`,`universal_identity`,`k8s`,`jwt`.

##### AWS Example

```hcl
connection "akeyless" {
  plugin    = "akeyless/akeyless"
  access_type = "aws_iam"
  access_id = "p-xxxxxxxxxx"
}
```

## Get involved

- Open source: https://github.com/akeyless/steampipe-plugin-akeyless
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
