package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/jaavier/sider"
)

var server = echo.New()
// /etc/letsencrypt/csr
func main() {
	fmt.Println("Hola Mundo!")
	server.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello Gopher!")
	})
	server.GET("/expire/:key/:value/:ttl", func(c echo.Context) error {
		key := c.Param("key")
		value := c.Param("value")
		ttl := c.Param("ttl")
		sider.Set(key, value)
		return c.String(200, fmt.Sprintf("Adding key and expire: %s, %s at %s", key, value, ttl))
	})
	server.GET("/get-key/:key", func(c echo.Context) error {
		key := c.Param("key")
		if content, err := sider.Get(key); err != nil {
			return c.String(400, "Key not found.")
		} else {
			return c.String(200, content)
		}
	})
	server.Logger.Fatal(server.StartTLS(":1337", "cert.pem", "privkey.pem"))
}
