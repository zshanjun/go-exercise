package main

import (
	"net"
	"time"
	"fmt"
	"bufio"
	"encoding/binary"
	"io"
)

var connMap map[uint32]*net.TCPConn

func main() {
	// listen for telnet client
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8889")
	if err != nil {
		panic("addr invalid")
	}
	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	// connect to server
	tunnel, err := net.DialTimeout("tcp", "localhost:8888", time.Second)
	if err != nil {
		panic(err)
	}

	var i uint32 = 1
	// storage conn
	connMap = make(map[uint32]*net.TCPConn)
	go receive(tunnel.(*net.TCPConn))
	for {
		fmt.Println("receive from client")
		conn, err := ln.AcceptTCP()
		if err != nil {
			panic(err)
		}
		go handle(conn, tunnel.(*net.TCPConn), i)
		connMap[i] = conn
		i++
	}
}

type header struct {
	LinkId uint32
	Len uint32
}

func receive(tunnel *net.TCPConn) {
	for {
		fmt.Println("read from server")
		reader := bufio.NewReader(tunnel)
		var h header
		binary.Read(reader, binary.LittleEndian, &h)
		buf := make([]byte, h.Len)
		io.ReadFull(reader, buf)
		conn := connMap[h.LinkId]
		fmt.Println(h.LinkId)
		conn.Write(buf)
	}
}

func handle(conn *net.TCPConn, tunnel *net.TCPConn, i uint32) {
	fmt.Println("handle connect")
	writer := bufio.NewWriter(tunnel)
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			conn.Close()
			return
		}
		if err != nil {
			panic(err)
		}
		binary.Write(writer, binary.LittleEndian, &header{i, uint32(n)})
		writer.Write(buf[:n])
		writer.Flush()
		fmt.Println("write to server")
	}
}
