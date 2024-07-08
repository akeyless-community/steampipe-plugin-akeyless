# Table: akeyless_role

Access roles

## Examples

### List all Roles with complete information

```sql
select
  *
from
  akeyless_role;
```

### List all Roles by Name, Created Date, and Rules

```sql
select
  role_name,
  creation_date,
  rules
from
  akeyless_role;
```

### List all Roles created in the last 30 days

```sql
select
  *
from
  akeyless_role
where
 creation_date >= CURRENT_TIMESTAMP - INTERVAL '30 days'
```
