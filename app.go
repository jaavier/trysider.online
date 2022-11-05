package main

import (
	"api/blacklist"
	"api/offers"
	"api/queue"
	rooms "api/rooms"
	"api/shortener"

	"github.com/labstack/echo/v4"
)

var server = echo.New()

const MAX_FREE = 2
const MAX_VALUE_LENGTH_FREE = 40
const MAX_PASSWORD_LENGTH_FREE = 64
const MAX_KEY_LENGTH_FREE = 40
const MAX_RETRIES_FREE = 5

func main() {
	server.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello Gopher!")
	})

	server.GET("/sha256/:text", func(c echo.Context) error {
		var text string = c.Param("text")
		return c.String(200, generateSha256(text))
	})

	server.GET("/blacklist", blacklist.Get)
	server.POST("/blacklist", blacklist.Post)
	server.DELETE("/blacklist", blacklist.Delete)

	server.GET("/rooms", rooms.Get)
	server.POST("/rooms", rooms.Post)
	server.DELETE("/rooms", rooms.Delete)

	server.GET("/offers", offers.Get)
	server.POST("/offers", offers.Post)
	server.DELETE("/offers", offers.Delete)

	server.GET("/queue", queue.Get)
	server.POST("/queue", queue.Post)
	server.DELETE("/queue", queue.Delete)

	server.GET("/shortener", shortener.Get)
	server.POST("/shortener", shortener.Post)
	server.DELETE("/shortener", shortener.Delete)

	server.Logger.Fatal(server.Start(":1337"))
	// server.Logger.Fatal(server.StartTLS(":1337", "cert.pem", "privkey.pem"))
}
