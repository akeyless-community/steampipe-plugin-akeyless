# Table: akeyless_target
Akeyless targets

## Examples

### List all Targets with complete information

```sql
select
  *
from
  akeyless_target;
```

### List all Targets by Name, Created Date, and Type

```sql
select
  target_name,
  creation_date,
  target_type
from
  akeyless_target;
```

### List all Targets created in the last 30 days

```sql
select
  *
from
  akeyless_target
where
 creation_date >= CURRENT_TIMESTAMP - INTERVAL '30 days';
```
