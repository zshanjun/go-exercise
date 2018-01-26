package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8888")
	if err != nil {
		panic("addr invalid")
	}
	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	for {
		conn, _ := ln.AcceptTCP()
		buf := make([]byte, 1024)
		for {
			fmt.Println("read from conn")
			n, err := conn.Read(buf)
			if err != nil {
				panic(err)
			}
			conn.Write(buf[:n])
		}
	}
}
