package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/jaavier/sider"
	"github.com/labstack/echo/v4"
)

var server = echo.New()

const MAX_FREE = 2
const MAX_VALUE_LENGTH_FREE = 40
const MAX_PASSWORD_LENGTH_FREE = 64
const MAX_KEY_LENGTH_FREE = 40
const MAX_RETRIES_FREE = 5

type NewKey struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Password string `json:"password"`
	Ttl      int64  `json:"ttl"`
}

func main() {
	server.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello Gopher!")
	})

	server.POST("/create", func(c echo.Context) error {
		var ip string = c.RealIP()

		if isBanned(ip) {
			return c.String(400, "You're banned for 1 day!")
		}

		if !checkLimit(ip) {
			return c.String(400, fmt.Sprintf("You reached the limit of free keys (%d). Please delete one and try again", MAX_FREE))
		}

		var newKey NewKey
		json.NewDecoder(c.Request().Body).Decode(&newKey)

		if len(newKey.Key) == 0 || len(newKey.Value) == 0 {
			return c.String(400, "I need key and value")
		}

		if len(newKey.Key) > MAX_KEY_LENGTH_FREE {
			return c.String(400, fmt.Sprintf("Key cannot be larger than %d", MAX_KEY_LENGTH_FREE))
		}

		if len(newKey.Password) > MAX_PASSWORD_LENGTH_FREE {
			return c.String(400, fmt.Sprintf("Password cannot be larger than %d", MAX_PASSWORD_LENGTH_FREE))
		}

		if len(newKey.Value) > MAX_VALUE_LENGTH_FREE {
			return c.String(400, fmt.Sprintf("Value cannot be larger than %d", MAX_VALUE_LENGTH_FREE))
		}

		if _, err := sider.Set(newKey.Key, newKey.Value); err != nil {
			return c.String(400, "Error setting key")
		}

		var keyLimit = fmt.Sprintf("limit.%s", ip)

		if lastLimit, err := sider.Get(keyLimit); err != nil {
			return c.String(400, "Error getting last limit")
		} else {
			lastLimitValue, _ := strconv.Atoi(lastLimit)
			lastLimitValue++
			sider.Set(keyLimit, fmt.Sprintf("%d", lastLimitValue))
		}

		if len(newKey.Password) > 0 {
			sider.Set(fmt.Sprintf("password.%s", newKey.Key), newKey.Password)
		}

		sider.Set(newKey.Key, fmt.Sprintf("%s", newKey.Value))

		return c.String(200, "OK")
	})

	server.GET("/get-key/:key", func(c echo.Context) error {
		key := c.Param("key")
		findPassword, _ := sider.Get(fmt.Sprintf("password.%s", key))
		if len(findPassword) > 0 {
			return c.String(400, "Require password.")
		}
		if content, err := sider.Get(key); err != nil {
			return c.String(400, "Key not found.")
		} else {
			return c.String(200, content)
		}
	})

	server.GET("/get-key/:key/:password", func(c echo.Context) error {
		var ip string = c.RealIP()
		if isBanned(ip) {
			return c.String(400, "You're banned.")
		}
		key := c.Param("key")
		password := c.Param("password")
		findPassword, _ := sider.Get(fmt.Sprintf("password.%s", key))
		if len(findPassword) > 0 {
			if findPassword != password {
				var retries int = getRetries(ip)
				retries++
				var retriesString string = fmt.Sprintf("%d", retries)
				var retriesKey string = fmt.Sprintf("retries.%s", ip)
				sider.Set(retriesKey, retriesString)
				if retries > MAX_RETRIES_FREE {
					sider.Set(fmt.Sprintf("banned.%s", ip), "true")
				}
				return c.String(400, "Invalid password.")
			}
		}
		if content, err := sider.Get(key); err != nil {
			return c.String(400, "Key not found.")
		} else {
			return c.String(200, content)
		}
	})
	server.Logger.Fatal(server.Start(":1337"))
	// server.Logger.Fatal(server.StartTLS(":1337", "cert.pem", "privkey.pem"))
}
