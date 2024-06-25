LOCAL_PATH=~/.steampipe/plugins/local/akeyless

install:
	go build -o ~/.steampipe/plugins/hub.steampipe.io/plugins/akeyless/akeyless@latest/steampipe-plugin-akeyless.plugin *.go

local:
	rm -f $(LOCAL_PATH)/akeyless.plugin
	go build -o $(LOCAL_PATH)/akeyless.plugin

