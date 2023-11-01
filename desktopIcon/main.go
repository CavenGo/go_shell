package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
)

func main() {
	a, _ := ReadLnk()
	fmt.Println(a)
}

// $shell = New-Object -ComObject WScript.Shell
// $shortcut = $shell.CreateShortcut("C:\路径\到\你的\快捷方式.lnk")
// $shortcut.TargetPath
// creat lnk : mklink /D "目标路径" "快捷方式路径"
// read lnk
func ReadLnk() ([]string, error) {
	desktopPath := `C:\Users\Artist\Desktop`

	lnklist := []string{}
	files, err := ioutil.ReadDir(desktopPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".lnk" {
			linkName := file.Name()
			lnklist = append(lnklist, linkName)
		}
	}
	//return lnklist, nil
	//通过lnklist获取所有lnk文件指向的地址
	realaddr := []string{}
	for _, vaule := range lnklist {
		var err error
		cmd3 := exec.Command("powershell", "-Command", `
$shell = New-Object -ComObject WScript.Shell
$shortcut = $shell.CreateShortcut("C:\Users\Artist\Desktop\`+vaule+`")
$shortcut.TargetPath
`)
		output, _ := cmd3.CombinedOutput()
		if err != nil {
			fmt.Println("获取"+vaule+"指向地址出错:", err)
		}
		realaddr = append(realaddr, string(output))
	}
	return realaddr, nil
}

// 获取原本的地址
func RealAddr(lnk string) string {
	var err error
	cmd3 := exec.Command("powershell", "-Command", `
$shell = New-Object -ComObject WScript.Shell
$shortcut = $shell.CreateShortcut("C:\Users\Artist\Desktop\Whale.exe.lnk")
$shortcut.TargetPath
`)
	output, _ := cmd3.CombinedOutput()
	if err != nil {
		fmt.Println("命令3执行出错:", err)
		return ""
	}
	return (string(output))
}
