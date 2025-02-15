package main

import (
	"fmt"
	"log"
	"net"
)

// handleConnection 处理每个客户端连接
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 读取客户端发送的数据
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Println("读取数据时出错:", err)
		return
	}

	// 打印接收到的数据
	message := string(buffer[:n])
	fmt.Println("接收到客户端消息:", message)

	// 向客户端发送响应
	response := "Pong"
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Println("发送响应时出错:", err)
		return
	}
}

func main() {
	// 监听本地的8080端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("监听端口时出错:", err)
	}
	defer listener.Close()

	fmt.Println("Ping服务已启动，监听端口: 8080")

	for {
		// 接受客户端连接
		conn, err := listener.Accept()
		if err != nil {
			log.Println("接受连接时出错:", err)
			continue
		}

		// 处理客户端连接
		go handleConnection(conn)
	}
}
