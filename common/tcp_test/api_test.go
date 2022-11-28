package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"testing"
)

func Benchmark_client_loop(b *testing.B) {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	// 输出当前建Dial函数的返回值类型, 属于*net.TCPConn类型
	fmt.Printf("客户端: %T\n", conn)
	if err != nil {
		// 连接的时候出现错误
		fmt.Println("err :", err)
		return
	}
	// 当函数返回的时候关闭连接
	defer conn.Close()
	// 获取一个标准输入的*Reader结构体指针类型的变量
	bt := []byte("ng-你好")
	ktds := 100
	for i := 0; i < ktds; i++ {
		conn.Write(append(bt, []byte(strconv.Itoa(i)+"\n")...))
	}
}

func Benchmark_client_go(b *testing.B) {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	// 输出当前建Dial函数的返回值类型, 属于*net.TCPConn类型
	fmt.Printf("客户端: %T\n", conn)
	if err != nil {
		// 连接的时候出现错误
		fmt.Println("err :", err)
		return
	}
	// 当函数返回的时候关闭连接
	defer conn.Close()
	// 获取一个标准输入的*Reader结构体指针类型的变量
	bt := []byte("go-你好")
	wg := sync.WaitGroup{}
	ktds := 500
	for i := 0; i < ktds; i++ {
		wg.Add(1)
		go func(j int) {
			conn.Write(append(bt, []byte(strconv.Itoa(j)+"\n")...))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
