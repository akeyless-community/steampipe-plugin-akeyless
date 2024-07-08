# Table: akeyless_gateway

Akeyless gateways

## Examples

### Basic Info

```sql
select
  gateway_id,
  display_name,
  cluster_name,
  cluster_url,
  customer_fragments,
  status,
  allowed,
  allowed_access_ids,
  default_protection_key_id,
  default_secret_location 
from
  akeyless_gateway;
```
