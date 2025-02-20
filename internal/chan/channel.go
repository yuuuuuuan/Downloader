package channel

import (
	"log"

	"github.com/yuuuuuuan/Downloader/internal/network"
)

func SendPostRequset(url string, reqChan chan map[string]interface{}, respChan chan map[string]interface{}) {
	for req := range reqChan {
		resp, err := network.HttpPost(url, req)
		if err != nil {
			//print("chan err1")
			log.Fatal(err)
		}
		res, err := network.ConvertRespToJson(resp)
		if err != nil {
			//print("chan err2")
			log.Fatal(err)
		}
		respChan <- res
	}
}
