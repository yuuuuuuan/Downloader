package main

import "fmt"

func init() {

}

func main() {
	// 创建 Downloader 实例
	d := &Downloader{
		URL:      "https://example.com/file.zip",
		FilePath: "/path/to/save/file.zip",
	}

	// 调用 StartDownload 方法
	err := d.StartDownload()
	if err != nil {
		fmt.Println("Download failed:", err)
	}
}
