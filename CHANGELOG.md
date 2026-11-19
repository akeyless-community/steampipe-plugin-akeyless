# v0.0.10

## Dependencies

- Recompiled with Go 1.26.11.
- Upgraded `steampipe-plugin-sdk` to v5.14.1 and transitive dependencies (`golang.org/x/crypto`, `golang.org/x/net`, `golang.org/x/oauth2`, `google.golang.org/grpc`, `go.opentelemetry.io/otel`, `hashicorp/go-getter`, `klauspost/compress`, `ulikunitz/xz`, `golang-jwt/jwt/v5`, `golang/glog`, `azidentity`) to remediate critical, high, medium and low vulnerabilities.

# v0.0.1

## First release

- New tables added
  - [akeyless_auth_method](https://hub.steampipe.io/plugins/akeyless-community/akeyless/tables/akeyless_auth_method)
  - [akeyless_gateway](https://hub.steampipe.io/plugins/akeyless-community/akeyless/tables/akeyless_gateway)
  - [akeyless_role](https://hub.steampipe.io/plugins/akeyless-community/akeyless/tables/akeyless_role)
  - [akeyless_item](https://hub.steampipe.io/plugins/akeyless-community/akeyless/tables/akeyless_item)
  - [akeyless_target](https://hub.steampipe.io/plugins/akeyless-community/akeyless/tables/akeyless_target)
