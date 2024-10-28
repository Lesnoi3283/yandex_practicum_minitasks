package main

import (
	"log"
	"net"
)

const (
	Port   = ":52001" // порт сервера
	MaxLen = 1024     // максимальный размер слайса
)

// handleConn обрабатывает запросы и вычисляет среднее арифметическое.
func handleConn(c net.Conn) {
	// допишите код
	// ...

	defer c.Close()

	for {
		buffer := make([]byte, MaxLen)
		n, err := c.Read(buffer)
		if err != nil {
			log.Printf("err while reading TCP request, err: %v", err)
			break
		}

		sum := 0
		for i := 0; i < len(buffer); i++ {
			sum += int(buffer[i])
		}
		sum /= n

		_, err = c.Write([]byte{byte(sum)})
		if err != nil {
			log.Printf("err while writing to TCP, err: %v", err)
			break
		}
	}
}

// TCPServer запускает сервер и ожидает соединений.
func TCPServer(addr *net.TCPAddr) {
	// допишите код
	// ...
	l, err := net.ListenTCP(addr.Network(), addr)
	if err != nil {
		log.Fatalf("cant start TCP server, err: %v", err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatalf("cant accept connection, err: %v", err)
		}
		go handleConn(c)
	}
}
