package pkg
import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"MineMonitor/internal/pkg/types"
	"MineMonitor/internal/utils"
)

func RunCrawler(saveDirectory string, ip_list []string){
	for _, iL := range ip_list {
    // array of open server
		var openServers []types.Server

    // generate ip
		generatedIP := IpGenerator(iL)
    log.Printf("IP generated with: %s", iL)
		time.Sleep(10 * time.Second)

		// result channel that stores all result
		resultsChan := make(chan types.Server, len(generatedIP))

		// limit / channel
		maxGoroutines := 2000
		limitChan := make(chan struct{}, maxGoroutines)

    // Channel for goroutine counter
    stopChan := make(chan struct{})
    
    // Goroutine counter
		go func() {
			for {
        select{
        case <- stopChan:
          return
        default:
          log.Printf("Current Goroutines: %v", len(limitChan))
          time.Sleep(1 * time.Second)
        }
			}
		}()

		// WaitGroup
		var wg sync.WaitGroup

    // for with generatedIP
		for _, v := range generatedIP {
			wg.Add(1)

      // Add empty struct to limitChan
      // It waits if there's no empty slot
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

        // Check ip
        // return: pkg.Server
				r, err := Status(ip, 25565, 2000)
				if err == nil && r.Version != "" {
					resultsChan <- r
				}

			}(v)
		}

    wg.Wait()
    close(resultsChan)

    // Stop Goroutine counter
    stopChan <- struct{}{}

    // Push each result to an Array
		for r := range resultsChan {
			openServers = append(openServers, r)
		}

    // Write result as .txt file in a directory
		utils.WriteResult(openServers, saveDirectory)
	}
}

