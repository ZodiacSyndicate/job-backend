module job-backend

go 1.12

replace (
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 => github.com/golang/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c => github.com/golang/crypto v0.0.0-20190325154230-a5d413f7728c
	golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a => github.com/golang/sys v0.0.0-20190215142949-d0b11bdaac8a
	golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223 => github.com/golang/sys v0.0.0-20190222072716-a9d3bda3a223
)

require (
	github.com/google/uuid v1.1.1
	github.com/jinzhu/gorm v1.9.5
	github.com/kr/pretty v0.1.0 // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo/v4 v4.1.5
	github.com/satori/go.uuid v1.2.0
)
