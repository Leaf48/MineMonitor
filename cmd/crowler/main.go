package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
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
	saveDirectory := constructer()

	ip_list := utils.Config().Ips

	for _, iL := range ip_list {
		var openServers []pkg.Server

		generatedIP := pkg.IpGenerator(iL)
		log.Print("IP generated: ")
		log.Println(iL)
		time.Sleep(10 * time.Second)

		// result / channel
		resultsChan := make(chan pkg.Server, len(generatedIP))

		// make waitGroup
		var wg sync.WaitGroup

		// limit / channel
		maxGoroutines := 2000
		limitChan := make(chan struct{}, maxGoroutines)

		go func() {
			for {
				fmt.Printf("Current Goroutines: ")
				fmt.Println(len(limitChan))
				time.Sleep(1 * time.Second)
			}
		}()

		for _, v := range generatedIP {
			wg.Add(1)

			limitChan <- struct{}{}

			go func(ip string) {
				// fmt.Println(ip)
				splitIp := strings.Split(ip, ".")[3]

				defer func() {
					<-limitChan
					wg.Done()
					if splitIp == "0" {
						fmt.Println(ip)
					}
				}()

				r, err := pkg.Status(ip, 25565, 2000)
				if err == nil && r.Version != "" {
					resultsChan <- r
				}

			}(v)
		}

		go func() {
			wg.Wait()
			close(resultsChan)
		}()

		for r := range resultsChan {
			openServers = append(openServers, r)
		}

		utils.WriteResult(openServers, saveDirectory)

	}
}
