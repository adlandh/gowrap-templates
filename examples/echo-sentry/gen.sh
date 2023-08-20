#!/bin/sh

oapi-codegen -old-config-style -generate types,server -o "openapi_gen.go" -package "main" "api.yaml"
gowrap gen -i ServerInterface -t https://raw.githubusercontent.com/adlandh/gowrap-templates/main/echo-sentry.gotmpl -o openapi_sentry_gen.go
gowrap gen -i AppInterface -t https://raw.githubusercontent.com/adlandh/gowrap-templates/main/sentry.gotmpl -o aplication_sentry_gen.go
