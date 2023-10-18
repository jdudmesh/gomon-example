module gohtmx

go 1.20

require (
	github.com/a-h/templ v0.2.408
	github.com/jmoiron/sqlx v1.3.5
	github.com/mattn/go-sqlite3 v1.14.17
	github.com/sirupsen/logrus v1.9.0
)

require (
	github.com/stretchr/testify v1.8.1 // indirect
	golang.org/x/sys v0.12.0 // indirect
)

replace github.com/jdudmesh/gomon-client => ../gomon-client
