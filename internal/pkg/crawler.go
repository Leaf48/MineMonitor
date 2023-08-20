package pkg
import (
	"fmt"
	"log"
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
		maxGoroutines := 1000
		limitChan := make(chan struct{}, maxGoroutines)

    // Channel for internet outage
    outageChan := make(chan struct{}, maxGoroutines)

    // Channel for when Goroutine stops
    stopChan := make(chan struct{})

    // Counter
    go func() {
      for {
        select{
        case <- stopChan:
          log.Println("Goroutine task finished!")
          return
        default:
          if !utils.IsConnected{
            log.Println("Outage has been detected. Reconnecting...")
            log.Printf("Current Outage: %v", len(outageChan))
            time.Sleep(3 * time.Second)

            if utils.IsConnected{
              log.Println("Outage has been solved. Starting...")
            }
          }else{
            fmt.Printf("Current: %v\n", len(limitChan) - len(outageChan))
            time.Sleep(2 * time.Second)
          }
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

      // Goroutine for checking server
			go func(ip string) {
				// splitIp := strings.Split(ip, ".")[3]

        // In case of outage
        if !utils.IsConnected{
          outageChan <- struct{}{}

          for !utils.IsConnected{
            time.Sleep(1 * time.Second)
          }
          <- outageChan
        }else{
          // Check ip
          // return: pkg.Server
          r, err := Status(ip, 25565, 4000)
          if err == nil && r.Version != "" {
            resultsChan <- r
          }
        }

        // Done Goroutine
				defer func() {
					<-limitChan
					wg.Done()
					// if splitIp == "0" {
					// 	fmt.Println(ip)
					// }
				}()
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
		utils.WriteResult(openServers, saveDirectory, iL)
	}
}

