package main

import (
	"fmt"
	"strconv"

	"github.com/jaavier/sider"
)

func getRetries(ip string) int {
	value, _ := sider.Get(fmt.Sprintf("retries.%s", ip))
	var intValue int = 0
	if len(value) > 0 {
		intValue, _ = strconv.Atoi(value)
	}
	return intValue
}

func isBanned(ip string) bool {
	value, _ := sider.Get(fmt.Sprintf("banned.%s", ip))
	return len(value) > 0
}

func checkLimit(ip string) bool {
	limit := fmt.Sprintf("limit.%s", ip)
	value, _ := sider.Get(limit)
	if len(value) == 0 {
		sider.Set(limit, "0")
		return true
	}
	toInt, _ := strconv.Atoi(value)
	return toInt < MAX_FREE
}
