package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	os := runtime.GOOS
	fmt.Println("当前操作系统：", os)
	if os == "windows" {
		Windwos()
	} else if os == "linux" {
		Linux()
	} else {
		fmt.Println("本命令不支持当前当前操作系统")
	}
}
func Windwos() {
	args := os.Args[1:] // 排除第一个参数，它是可执行文件的路径
	if len(args) > 0 {
		// 读取其他参数
		if len(args) == 1 && args[0] == "?" {
			fmt.Println("运行godinstall不带参数自动安装curl，wget,jq,wsl工具")
			fmt.Println("winget install toolname 使用winget指定安装curl，wget,jq工具")
			fmt.Println("curl -o (filepath) url 使用curl指定安装go,gcc/g++,cmake")
			return
		}
		if len(args) == 2 {

		}
	} else {
		tool := []string{"curl", "jq", "wget", "wsl"}
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

func Linux() {
	args := os.Args[1:] // 排除第一个参数，它是可执行文件的路径
	if len(args) > 0 {
		// 读取其他参数
		if len(args) == 1 && args[0] == "?" {
			fmt.Println("运行godinstall不带参数自动安装curl，wget,jq工具")
			fmt.Println("winget install toolname 使用winget指定安装curl，wget,jq工具")
			fmt.Println("curl -o (filepath) url 使用curl指定安装go,gcc/g++,cmake")
			return
		}
		if len(args) == 2 {

		}
	} else {
		cmd := exec.Command("sudo", "apt", "update")
		cmd.Run()
		tool := []string{"vim", "jq", "net-tools", "netplan.io", "lshw"}
		for _, _ = range tool {

		}
		cmd = exec.Command("bash", "-c", "mkdir /root/source")
		cmd.Run()
		cmd = exec.Command("bash", "-c", "which go")
		goin, _ := cmd.Output()
		if string(goin) != "" {
			cmd = exec.Command("bash", "-c", "cp -r ./source/golang /root")
			cmd.Run()
		}
	}
}
