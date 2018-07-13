package lib

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const DownloadsDir = "./"

var Finshed bool

func DownloadFile(url, dir string, pool *GoPool) {
	defer func() {
		go CountDownloadedFile.Add()
		pool.Done()
	}()
	stringArray := strings.Split(url, "/")
	fileName := stringArray[len(stringArray)-1]
	os.MkdirAll(DownloadsDir+dir, os.ModePerm)
	out, err := os.Create(DownloadsDir + dir + "/" + fileName)
	if err != nil {
		fmt.Println(url, err)
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf(url+": bad status: %s\n", resp.Status)
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println(url, err)
	}
}

// concurrent 下载并发数
func DownloadSeriesFile(series []Series, dir string, concurrent int) {
	goPool := NewInstance(concurrent)
	// 下载文件
	fmt.Println("Downloading...")
	for _, serie := range series {
		for _, instance := range serie.InstanceList {
			goPool.Add()
			go DownloadFile(instance.ImageId, dir, goPool)
		}
	}
	goPool.Wait()
	Finshed = true
	time.Sleep(time.Hour)
}

// CountSeriesFile
func CountSeriesFile(series []Series) int {
	count := 0
	for _, serie := range series {
		count += len(serie.InstanceList)
	}
	return count
}
