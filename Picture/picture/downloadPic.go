package picture

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func downloadImage(imageURL string, savePath string, proxyserve string) error {
	err := download_With_NOproxy(imageURL, savePath, proxyserve)
	if err != nil {
		proxyURL, err := url.Parse(proxyserve)
		if err != nil {
			return fmt.Errorf("无效的代理地址: %v", err)
		}

		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}

		httpClient := &http.Client{
			Transport: transport,
			Timeout:   time.Second * 2,
		}

		response, err := httpClient.Get(imageURL)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			return fmt.Errorf("HTTP请求失败，状态码：%d", response.StatusCode)
		}

		fileName := path.Base(imageURL)
		filePath := savePath + "\\" + fileName

		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, response.Body)
		if err != nil {
			return err
		}

		fmt.Printf("图像已下载：%s\n", filePath)
		return nil
	}
	return nil
}

func download_With_NOproxy(imageURL string, savePath string, proxyserve string) error {
	client := http.Client{
		Timeout: time.Second * 2,
	}
	response, err := client.Get(imageURL)
	if err != nil {
		fmt.Println("下载图片失败，准备使用代理")
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("下载图片失败，准备使用代理")
		return fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	fileName := path.Base(imageURL)
	filePath := savePath + "\\" + fileName

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("下载图片失败，准备使用代理")
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("下载图片失败，准备使用代理")
		return err
	}

	fmt.Printf("Image downloaded: %s\n", filePath)
	return nil
}

// 从数据向前url中解析图片：<img src="data:image/jpeg;base64,/9j.ooooAKKKKACiiigD/9kHBwcHBwcH" alt="Embedded Image">
func DownloadImageFromDataURL(dataURL, savePath, proxyURL string, index int) error {
	err := downloadImageFromDataURL_with_Noproxyserver(dataURL, savePath, index)
	if err != nil {
		downloadImageFromDataURL_with_proxyserver(dataURL, savePath, proxyURL, index)
	}
	return nil
}

func downloadImageFromDataURL_with_Noproxyserver(dataURL, savePath string, index int) error {
	currentTime := time.Now()                         // 获取当前时间
	time := currentTime.Format("2006-01-02-15-04-05") // 将时间格式化为字符串
	// 解码数据URL
	decoded, err := base64.StdEncoding.DecodeString(strings.Split(dataURL, ",")[1])
	if err != nil {
		return err
	}

	in := strconv.Itoa(index)
	// 创建图像文件
	file, err := os.Create(savePath + "\\" + time + "_" + in + ".jpg")
	if err != nil {
		return err
	}
	defer file.Close()

	// 将解码后的图像数据写入文件
	_, err = file.Write(decoded)
	if err != nil {
		return err
	}

	fmt.Printf("图像已下载：%s\n", savePath+"\\"+time+".jpg")
	return nil
}

func downloadImageFromDataURL_with_proxyserver(dataURL, savePath string, proxyURL string, index int) error {
	// 创建代理客户端
	currentTime := time.Now()                         // 获取当前时间
	time := currentTime.Format("2006-01-02-15-04-05") // 将时间格式化为字符串
	proxyURLParsed, err := url.Parse(proxyURL)
	if err != nil {
		return fmt.Errorf("无效的代理地址: %v", err)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURLParsed),
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	// 创建请求
	request, err := http.NewRequest("GET", dataURL, nil)
	if err != nil {
		return err
	}

	// 发送请求
	response, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP请求失败，状态码：%d", response.StatusCode)
	}

	in := strconv.Itoa(index)
	// 创建图像文件
	file, err := os.Create(savePath + "\\" + time + "_" + in + ".jpg")
	if err != nil {
		return err
	}
	defer file.Close()

	// 将响应体中的图像数据写入文件
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("图像已下载：%s\n", savePath+"\\"+time+".jpg")
	return nil
}
