package main

import (
	"bufio"
	"log"
	"net"
	"socks/protocol"
)

func handler(conn net.Conn) {
	defer conn.Close()
	// 握手阶段
	reader := bufio.NewReader(conn)
	buf := make([]byte, 1+1+255)
	n, err := reader.Read(buf)
	if err != nil {
		log.Printf("read failed: %v", err)
		return
	}
	res, err := protocol.HandleHandshake(buf[:n])
	if err != nil {
		log.Printf("err happen: %v", err)
		return
	}
	conn.Write(res)
	log.Printf("handshark finish, send reply %v", res)
	// 连接阶段
	buf = make([]byte, 1+1+1+1+255+2)
	n, err = reader.Read(buf)
	if err != nil {
		log.Printf("read failed: %v\n", err)
		return
	}
	res, err = protocol.HandleConnect(buf[:n])
	if err != nil {
		log.Printf("err hapened: %v\n", err)
		return
	}
	conn.Write(res)
}

func main() {
	listen, err := net.Listen("tcp", ":9999")
	log.Printf("socks5 server started, listening at port 9999")
	if err != nil {
		log.Printf("Listen failed: %v\n", err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("error happen: %v", err)
			continue
		}
		go handler(conn)
	}
}
