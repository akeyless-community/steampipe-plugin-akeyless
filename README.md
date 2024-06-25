![image]()

# Akeyless Plugin for Steampipe

Use SQL to query items, auth mehtods, roles, targets, gateways and more from Akeyless.

* **[Get started →](https://hub.steampipe.io/plugins/akeyless/akeyless)**
* Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/akeyless/akeyless/tables)
* Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
* Get involved: [Issues](https://github.com/akeyless/steampipe-plugin-akeyless/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io/downloads):

```shell
steampipe plugin install akeyless/akeyless
```

[Configure the plugin](https://hub.steampipe.io/plugins/akeyless/akeyless#configuration) with your authentication method using the configuration file:

```shell
vi ~/.steampipe/config/akeyless.spc
```

Start Steampipe:

```shell
steampipe query
```

Run a sample query:

```sql
select
  role_name
from
  akeyless_role;
```

Or run the query from your command line without starting Steampipe:

```shell
steampipe query "select role_name from akeyless_role;"
```

## Developing

Prerequisites:

* [Steampipe](https://steampipe.io/downloads)
* [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/akeyless-community/steampipe-plugin-akeyless.git
cd steampipe-plugin-akeyless
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```sh
make
```

Configure the plugin:

```sh
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/akeyless.spc
```

Try it!

```shell
steampipe query
> .inspect akeyless
```

Further reading:

* [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
* [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing
