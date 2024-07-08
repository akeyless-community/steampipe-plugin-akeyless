# Table: akeyless_item

Akeyless items

## Examples

### List all Items with complete information

```sql
select
  *
from
  akeyless_item;
```

### List all Items by Name and Type

```sql
select
  item_name,
  item_sub_type
from
  akeyless_item;
```
### List all Items created in the last 30 days

```sql
select
  *
from
  akeyless_item
where
 creation_date >= CURRENT_TIMESTAMP - INTERVAL '30 days'
```
