package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"syscall"
)

func handle(conn net.Conn) {
	fmt.Println(fmt.Sprintf("Connection from %v", (conn.RemoteAddr())))
	defer conn.Close()

	os.Mkdir("/tmp/sea", os.ModeDir)
	filename := fmt.Sprintf("/tmp/sea/%v", conn.RemoteAddr())

	out, err := os.Create(filename + ".out")
	if err != nil {
		panic(err)
	}
	go io.Copy(out, conn)

	syscall.Mkfifo(filename, 0666)
	defer os.Remove(filename)

	for {
		fifo, err := os.OpenFile(filename, os.O_RDONLY, 0666)
		if err != nil {
			break
		}
		io.Copy(conn, fifo)
		fifo.Close()
	}
}

func splash() {
	fmt.Printf(" ▞▀▘▞▀▖▝▀▖\n ▝▀▖▛▀ ▞▀▌\n ▀▀ ▝▀▘▝▀▘\n ")

}

func main() {
	splash()
	ln, err := net.Listen("tcp", ":1337")
	fmt.Printf("listening on 1337...\n")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}
