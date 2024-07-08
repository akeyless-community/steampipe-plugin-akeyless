# Table: akeyless_auth_method

Akeyless authentication methods

## Examples

### List all Authentication Methods with complete information

```sql
select
  *
from
  akeyless_auth_method;
```

### List all Authentication Methods by Name and Date Created

```sql
select
  auth_method_name,
  creation_date,
from
  akeyless_auth_method;
```
