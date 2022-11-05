package chat

import "github.com/labstack/echo/v4"

func Get(c echo.Context) error {
	return c.String(200, "GET ROOM")
}

func Post(c echo.Context) error {
	return c.String(200, "CREATE ROOM")
}

func MessagePost(c echo.Context) error {
	return c.String(200, "POST MESSAGE IN ROOM")
}

func Delete(c echo.Context) error {
	return c.String(200, "DELETE ROOM")
}
