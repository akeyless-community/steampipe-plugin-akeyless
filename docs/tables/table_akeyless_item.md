# Table: akeyless_item

Akeyless items

## Examples

### Basic Info

```sql
select
  item_id,
  account_id,
  creation_date,
  modification_date,
  item_name,
  item_sub_type,
  last_version,
  with_customer_fragment,
  is_enabled,
  protection_key_name,
  client_permissions,
  item_state,
  rotation_interval,
  item_general_info,
  item_targets_assoc,
  delete_protection,
  is_access_request_enabled,
  access_request_status,
  next_rotation_date,
  last_rotation_date,
  auto_rotate 
from
  akeyless_item;
```
