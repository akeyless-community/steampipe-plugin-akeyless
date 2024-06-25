# Table: akeyless_auth_method

Akeyless authentication methods

## Examples
### Basic Info

```sql
select
    auth_method_name,
    creation_date,
    modification_date,
    access_date,
    account_id,
    ttl,
    rules_type,
    force_sub_claims,
    access_info
from
    akeyless_auth_method;
```
