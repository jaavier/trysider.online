package blacklist

import "github.com/jaavier/sider"

func addIp(ip string) bool {
	if isIpBanned(ip) {
		return true
	}
	result, _ := sider.RPush("banned:ip", ip)
	return result
}

func isIpBanned(ip string) bool {
	var _, err = sider.IndexOf("banned:ip", ip)
	return err == nil
}

func removeIp(ip string) bool {
	if !isIpBanned(ip) {
		return false
	}
	if sider.DeleteItemByContent("banned:ip", ip) {
		return true
	}
	return false
}
