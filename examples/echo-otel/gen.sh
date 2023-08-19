#!/bin/sh

oapi-codegen -old-config-style -generate types,server -o "openapi_gen.go" -package "main" "api.yaml"
gowrap gen -i ServerInterface -t https://raw.githubusercontent.com/adlandh/gowrap-templates/main/echo-otel.gotmpl -o openapi_otel_gen.go
