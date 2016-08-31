package main

import (
	"fmt"
	"net"
)

func main() {
    fmt.Println("Start a TCP server ... ")

    listener, err := net.Listen("tcp", ":12345")

    if err != nil {
        fmt.Println("Error ", err.Error())
        return;
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            return;
        }

        go doStuff(conn)
    }
}

func doStuff(conn net.Conn) {
    for {
        buf := make([]byte, 512)
        _, err := conn.Read(buf)

        if err != nil {
            return;
        }

        fmt.Println("Got message : ", string(buf));
        
        hello := []byte{'H', 'i', '\n'}
        
        conn.Write(hello)

    }
}