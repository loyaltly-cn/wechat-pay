package io

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile() (map[string]interface{}, error) {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("无法打开 JSON 配置文件:", err)
	}
	defer file.Close()

	var config map[string]interface{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("无法解码 JSON 文件:", err)
	}

	return config, err
}
