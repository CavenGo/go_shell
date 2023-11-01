package main

import "web/picture"

func main() {
	websiteURL := "https://zztt15.com/archives/16378.html"
	savePath := "G:\\sR\\1抖音\\郑一亿\\p" // 保存图片的路径

	picture.GetHTMLcode(websiteURL, savePath)
}
