package main

import (
	"fmt"

	pkg "MineMonitor/internal/pkg"
)

func main() {

	r, err := pkg.Status("hypixel.net", 25565, 500)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.Port)

}
