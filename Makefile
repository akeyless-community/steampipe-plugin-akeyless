INSTALL_PATH=~/.steampipe/plugins/hub.steampipe.io/plugins/akeyless-community/akeyless@latest
LOCAL_PATH=~/.steampipe/plugins/local/akeyless

install:
	go build -o $(INSTALL_PATH)/steampipe-plugin-akeyless.plugin

local:
	go build -o $(LOCAL_PATH)/akeyless.plugin
