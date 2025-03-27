package downloader

import (
	"fmt"

	"github.com/yuuuuuuan/Downloader/storage"
)

type Downloader struct {
	URL      string
	FilePath string
}

func (d *Downloader) StartDownload() error {
	fmt.Println("Starting download:", d.URL)
	// 使用网络模块下载文件
	fileData, err := network.Post(d.URL)
	if err != nil {
		return err
	}
	// 使用存储模块保存文件
	err = storage.SaveFile(d.FilePath, fileData)
	if err != nil {
		return err
	}
	fmt.Println("Download complete!")
	return nil
}
