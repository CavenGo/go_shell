package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	args := os.Args[1:] // 排除第一个参数，它是可执行文件的路径
	if len(args) > 0 {
		// 读取其他参数
		if len(args) == 1 && args[0] == "?" {
			fmt.Println("本命令只在windows操作系统下可执行")
			fmt.Println("运行godinstall不带参数自动安装curl，wget,jq工具")
			fmt.Println("winget install toolname 使用winget指定安装curl，wget,jq工具")
			fmt.Println("curl -o (filepath) url 使用curl指定安装go,gcc/g++,cmake")
			return
		}
		if len(args) == 2 {

		}
	} else {
		tool := []string{"curl", "jq", "wget"}
		// 使用winget可以安装的命令
		for _, vaule := range tool {
			_, err := exec.LookPath(vaule)
			if err != nil {
				fmt.Println(vaule + " 命令不存在，正在安装...")
				// 使用 winget 命令安装 curl
				cmd := exec.Command("winget", "install", vaule)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					fmt.Println("安装 "+vaule+" 失败:", err)
					return
				}
				fmt.Println(vaule + " 安装成功！")
			} else {
				fmt.Println(vaule + " 工具已存在")
			}
		}
	}
}
