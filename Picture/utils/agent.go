package utils

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
)

// 打开 HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Internet Settings 的proxyServer键来查看当前系统是否使用代理
func AgentAddr() string {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.QUERY_VALUE)
	if err != nil {
		fmt.Println("无法打开注册表项:", err)
		return ""
	}
	defer key.Close()

	proxyEnabled, _, err := key.GetIntegerValue("ProxyEnable")
	if err != nil {
		fmt.Println("无法读取代理启用状态:", err)
		return ""
	}

	var proxyaddr string

	if proxyEnabled != 0 {
		proxyServer, _, err := key.GetStringValue("ProxyServer")
		if err != nil {
			fmt.Println("无法读取代理服务器地址:", err)
			return ""
		}
		fmt.Println("当前系统启用了代理服务，代理地址:", proxyServer)
		proxyaddr = proxyServer
	} else {
		fmt.Println("当前系统未启用代理服务")
	}
	return proxyaddr
}
