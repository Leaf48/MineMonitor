package main

import (
	"fmt"
	"os"
	"time"

	"MineMonitor/internal/pkg"
	"MineMonitor/internal/utils"
)

// Make default directory
func constructer() string {
	fileName := fmt.Sprintf("%d", time.Now().Unix())
	if err := os.Mkdir(fileName, 0755); err != nil {
		panic(err)
	}
	return fileName
}

func main() {
  // create directory to save .txt file
	saveDirectory := constructer()

  // get target ip
	ip_list := utils.Config().Ips

  pkg.RunCrawler(saveDirectory, ip_list)
}
