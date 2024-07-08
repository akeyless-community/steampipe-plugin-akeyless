# Table: akeyless_target
Akeyless targets
## Examples

### Basic Info

```sql
select
  target_name,
  creation_date,
  modification_date,
  access_date,
  target_type,
  with_customer_fragment,
  protection_key_name,
  client_permissions,
  last_version,
  attributes,
  is_access_request_enabled,
  access_request_status 
from
  akeyless_target;
```
