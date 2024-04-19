module github.com/adlandh/gowrap-templates/helpers/otel

go 1.22.2

replace github.com/adlandh/gowrap-templates/helpers => ../

require (
	github.com/adlandh/gowrap-templates/helpers v0.0.0-00010101000000-000000000000
	github.com/goccy/go-json v0.10.2
	github.com/labstack/echo/v4 v4.12.0
	go.opentelemetry.io/otel v1.25.0
	go.opentelemetry.io/otel/trace v1.25.0
)

require (
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
