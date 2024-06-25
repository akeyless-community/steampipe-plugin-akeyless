ARG GO_VERSION
FROM golang:1.22-alpine AS builder

WORKDIR /src

COPY . .

RUN go mod download

RUN GOOS=linux go build -o /build/akeyless.plugin

########################################################
# Plugin tests
FROM debian:12-slim AS plugintest

ARG ACCESS_ID
ARG ACCESS_KEY
ARG API_URL

RUN apt-get update -y && \
    apt-get install -y curl wget && \
    useradd -m steampipe && \
    /bin/sh -c "$(curl -fsSL https://steampipe.io/install/steampipe.sh)"

USER steampipe:steampipe

WORKDIR /home/steampipe

#init steampipe folders
RUN steampipe query "select 1;"

#Copy plugin as local
COPY --from=builder  /build/akeyless.plugin /home/steampipe/.steampipe/plugins/local/akeyless/akeyless.plugin

#Copy script to create plugin config
COPY   --chown=steampipe:steampipe create_plugin_config.sh  /home/steampipe/create_plugin_config.sh

#Run script to create  plugin config
RUN chmod +x create_plugin_config.sh && \
    ./create_plugin_config.sh $ACCESS_ID $ACCESS_KEY "/home/steampipe/.steampipe/config/akeyless.spc" $API_URL


#Test SQL queries
RUN steampipe query "SELECT role_name FROM akeyless.akeyless_role;" && sleep 1s && \
    steampipe query "SELECT item_name FROM akeyless.akeyless_item;" && sleep 1s && \
    steampipe query "SELECT auth_method_name FROM akeyless.akeyless_auth_method;" && sleep 1s && \
    steampipe query "SELECT target_name FROM akeyless.akeyless_target;" && sleep 1s && \
    steampipe query "SELECT cluster_name FROM akeyless.akeyless_gateway;" && sleep 1s

