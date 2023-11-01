package picture

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"golang.org/x/net/html"
	"io"
	"log"
	"strings"
	"time"
	"web/utils"
)

// 通过CHrome浏览器获取页面源码，所以需要安装CHrome页面源码
func GetHTMLcode(websiteURL, savePath string) {
	proxyServers := utils.AgentAddr()
	err := getHTML_with_NOproxy(websiteURL, savePath, proxyServers)
	if err != nil {
		if proxyServers == "" {
			fmt.Println("没有代理，g")
			return
		}
		opts := append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),
			chromedp.Flag("disable-gpu", true),
			chromedp.Flag("no-sandbox", true),
			chromedp.Flag("disable-dev-shm-usage", true),
			chromedp.ProxyServer(proxyServers),
		)

		ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
		defer cancel()

		ctx, cancel = chromedp.NewContext(ctx)
		defer cancel()

		var htmlContent string

		err := chromedp.Run(ctx,
			chromedp.Navigate(websiteURL),
			chromedp.OuterHTML("html", &htmlContent),
		)
		if err != nil {
			log.Fatal(err)
		}

		links, err := extractImageLinks(strings.NewReader(htmlContent), websiteURL)
		if err != nil {
			log.Fatal(err)
		}
		for index, link := range links {
			if strings.HasPrefix(link, "data") {
				err = DownloadImageFromDataURL(link, savePath, proxyServers, index)
				if err != nil {
					fmt.Printf("Failed to download image: %v\n", err)
				}
			} else if strings.HasPrefix(link, "http") {
				err = downloadImage(link, savePath, proxyServers)
				if err != nil {
					fmt.Printf("Failed to download image: %v\n", err)
				}
			}
		}
	}
}

func getHTML_with_NOproxy(websiteURL, savePath, proxyServers string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 设置超时时间
	ctx, cancel = context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var htmlContent string

	err := chromedp.Run(ctx,
		chromedp.Navigate(websiteURL),
		chromedp.OuterHTML("html", &htmlContent),
	)
	if err != nil {
		//log.Fatal(err)
		fmt.Println("访问网页失败，准备使用代理")
		return err
	}

	links, err := extractImageLinks(strings.NewReader(htmlContent), websiteURL)

	if err != nil {
		//log.Fatal(err)
		fmt.Println("访问网页失败，准备使用代理")
		return err
	}

	for index, link := range links {
		if strings.HasPrefix(link, "data") {
			err = DownloadImageFromDataURL(link, savePath, proxyServers, index)
			if err != nil {
				fmt.Printf("Failed to download image: %v\n", err)
			}
		} else if strings.HasPrefix(link, "http") {
			err = downloadImage(link, savePath, proxyServers)
			if err != nil {
				fmt.Printf("Failed to download image: %v\n", err)
			}
		}
	}
	return nil
}

func extractImageLinks(body io.Reader, baseURL string) ([]string, error) {
	var links []string

	tokenizer := html.NewTokenizer(body)

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			return nil, err
		}

		token := tokenizer.Token()

		if tokenType == html.StartTagToken && token.Data == "img" {
			for _, attr := range token.Attr {
				if attr.Key == "src" {
					imageURL := attr.Val
					if !strings.HasPrefix(imageURL, "http") && !strings.HasPrefix(imageURL, "data") {
						imageURL = baseURL + imageURL
					}
					links = append(links, imageURL)
					break
				}
			}
		}
	}

	return links, nil
}
