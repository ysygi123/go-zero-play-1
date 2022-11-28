package main

import (
	"fmt"
	"go-zero-play-1/common/tcp_test/ff"
	"net"
)

func main() {
	// 监听当前的tcp连接
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	fmt.Printf("服务端: %T=====\n", listen)
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		fmt.Println("当前建立了tcp连接")
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		// 对于每一个建立的tcp连接使用go关键字开启一个goroutine处理
		ff.Process(conn)
	}
}
