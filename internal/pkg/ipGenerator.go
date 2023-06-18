package pkg

import (
	"fmt"
	"strconv"
	"strings"
)

func IpGenerator(ip string) []string {
	ipAddr := strings.Split(ip, ".")
	firstIndex, _ := strconv.Atoi(ipAddr[0])
	secondIndex, _ := strconv.Atoi(ipAddr[1])

	var ips []string

	for b := 0; b <= 255; b++ {
		for c := 0; c <= 255; c++ {
			i := fmt.Sprintf(
				"%d.%d.%d.%d",
				firstIndex, secondIndex, b, c,
			)

			ips = append(ips, i)
		}
	}
	return ips
}
