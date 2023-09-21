module gohtmx

go 1.20

require github.com/labstack/echo/v4 v4.11.1

require (
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/james-barrow/golang-ipc v1.2.4 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/mod v0.10.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	golang.org/x/tools v0.9.1 // indirect
)

require github.com/jdudmesh/gomon/pkg/client v0.0.0

replace github.com/jdudmesh/gomon/pkg/client v0.0.0 => ../gomon-client
