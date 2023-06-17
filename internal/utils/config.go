package utils

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

type Sconfig struct {
	Ips []string
}

var (
	errConfigNotFound = errors.New("configが見つかりません")
)

func Config() Sconfig {
	var paths = []string{"", "../../", "../", "./"}
	config_str := Sconfig{}

	for _, v := range paths {
		dir := v + "checklist.yaml"

		readF, err := os.ReadFile(dir)

		if err == nil {
			yaml.Unmarshal(readF, &config_str)
			break
		}
	}
	return config_str
}
