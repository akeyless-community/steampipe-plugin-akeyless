#!/bin/sh

# Define input parameters with default values
ACCESS_ID=$1
ACCESS_KEY=$2
CONFIG_PATH=$3
API_URL=$4

# Create the configuration file
cat <<EOF > "${CONFIG_PATH}"
connection "akeyless" {
    plugin  = "local/akeyless"
    access_type = "api_key"
    access_id = "${ACCESS_ID}"
    access_key = "${ACCESS_KEY}"
    api_url = "${VAULT_URL}"
}
EOF

echo "Configuration file created at ${CONFIG_PATH}"
