package utils

import (
	"MineMonitor/internal/pkg"
	"fmt"
	"os"
	"time"
)

func WriteResult(servers []pkg.Server, saveDirectory string) {
	local := time.Now()
	fileName := fmt.Sprintf("%s/%d-%d-%d-%d-%d-%d.txt", saveDirectory, local.Year(), int(local.Month()), local.Day(), local.Hour(), local.Minute(), local.Second())

	// Create file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Open file
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != err {
		fmt.Println(err)
		return
	}

	// Write
	for i, v := range servers {

		newline := fmt.Sprintf("%d: %s", i, v.Ip)
		_, _ = fmt.Fprintln(f, newline)

		newline = fmt.Sprintf("	Version: %s", v.Version)
		_, _ = fmt.Fprintln(f, newline)

		newline = fmt.Sprintf("	Players: %s", v.Players)
		_, _ = fmt.Fprintln(f, newline)

		newline = fmt.Sprintf("	Player name: %s", v.PlayerList)
		_, _ = fmt.Fprintln(f, newline)

		newline = fmt.Sprintf("	MOTD: %s", v.Motd)
		_, _ = fmt.Fprintln(f, newline)

		_, _ = fmt.Fprintln(f, "")
	}
}
