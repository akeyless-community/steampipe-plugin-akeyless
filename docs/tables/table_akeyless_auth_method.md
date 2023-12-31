# Table: akeyless_auth_method

Each Authentication Method object is associated with an Access Role that grants permission (including Create, Read, Update, Delete, List, and Deny) to this Identity on Secrets, Targets, Roles, and Authentication Method objects stored inside the Akeyless SaaS solution.

## Examples

### Basic Info

```sql
select
   path,
   auth_method_access_id,
   access_info_rules_type,
   access_info_jwt_ttl 
from
   akeyless_auth_method;
```

### Lookup by Access ID

```sql
select
   path,
   auth_method_access_id,
   access_info_rules_type,
   access_info_jwt_ttl 
from
   akeyless_auth_method 
where
   auth_method_access_id = "p - 0sxjff4c8382";
```

### Count all Auth Methods by Type

```sql
select
   access_info_rules_type,
   count(*) as count_by_type 
from
   akeyless_auth_method 
group by
   access_info_rules_type;
```
