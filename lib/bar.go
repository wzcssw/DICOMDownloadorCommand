package lib

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

/*
	计算进度
*/

type Countor struct {
	// sync.Mutex
	sync.RWMutex
	Val int
}

var CountDownloadedFile = &Countor{}

func (countor *Countor) Add() {
	countor.RLock()
	defer countor.RUnlock()
	countor.Val++
}

// ShowProcess 显示进度条
func ShowProcess(count int) {
	time.Sleep(8e8)
	for {
		if Finshed {
			go func() {
				time.Sleep(1e9)
				fmt.Println()
				os.Exit(0)
			}()
		}
		per := float64(CountDownloadedFile.Val) / float64(count) * 100.0
		str := strconv.Itoa(CountDownloadedFile.Val) + "/" + strconv.Itoa(count) + " [" + bar(int(per)/2, 50) + "] " + strconv.Itoa(int(per)) + "%"
		fmt.Printf("\r%s", str)
		time.Sleep(3e8)
	}
}

func bar(count, size int) string {
	str := ""
	for i := 0; i < size; i++ {
		if i < count {
			str += "="
		} else {
			str += " "
		}
	}
	return str
}
