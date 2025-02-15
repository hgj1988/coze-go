package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// 服务器地址和端口
	serverAddr := "localhost:8080"

	// 连接到服务器
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatalf("无法连接到服务器: %v", err)
	}
	defer conn.Close()

	fmt.Println("已连接到服务器:", serverAddr)

	// 发送消息到服务器
	message := "Ping"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatalf("发送消息时出错: %v", err)
	}

	fmt.Println("已发送消息:", message)

	// 从服务器接收响应
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("接收响应时出错: %v", err)
	}

	fmt.Println("收到服务器响应:", response)
}
