package main

import (
	"log"
	"mermaid/boot"
)

func main() {
	if err := boot.Boot(); err != nil {
		log.Fatalf("程序启动失败,%s", err.Error())
	}
}
