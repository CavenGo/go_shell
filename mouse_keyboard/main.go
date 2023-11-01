package main

import (
	"github.com/go-vgo/robotgo"
)

func main() {
	//// 获取当前鼠标位置
	//x, y := robotgo.GetMousePos()
	//fmt.Println("当前鼠标位置：", x, y)

	// 移动鼠标到指定位置
	//robotgo.MoveMouseSmooth(100, 100, 1.0, 100.0)

	// 模拟鼠标左键点击
	robotgo.MouseClick("left", true)
	//robotgo.MouseClick("left", true)

	// 延迟一段时间
	robotgo.MilliSleep(500)

	// 模拟鼠标右键点击
	robotgo.MouseClick("right", true)
}
