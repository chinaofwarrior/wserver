package main

import (
	"fmt"
	"log"

	conf "github.com/social/wserver/config"
	conn "github.com/social/wserver/connection"
)

func main() {
	config, err := conf.LoadConfig("handlers.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	handlers := make(map[string]conn.CommandHandler)
	// 根据配置文件注册处理器
	for _, hc := range config.Handlers {
		handler, err := conn.HandlerFactory(hc.Name)
		if err != nil {
			log.Printf("Failed to create handler for command %s: %v", hc.Type, err)
			continue
		}
		handlers[hc.Name] = handler
	}

	// 创建命令通道
	commandChan := make(chan conn.Command)

	// 启动TCP服务器
	go conn.StartTCPServer(":8011", commandChan)

	// 处理命令
	for cmd := range commandChan {
		if handler, ok := handlers[cmd.Type]; ok {
			handler.Execute(cmd)
		} else {
			fmt.Printf("No handler registered for command type: %s\n", cmd.Type)
		}
	}
}
