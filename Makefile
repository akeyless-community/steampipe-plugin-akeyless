LOCAL_PATH=~/.steampipe/plugins/local/akeyless

install:
	go build -o ~/.steampipe/plugins/hub.steampipe.io/plugins/akeyless-community/akeyless@latest/steampipe-plugin-akeyless.plugin *.go

local:
	go build -o $(LOCAL_PATH)/akeyless.plugin
