package main

import (
	"log"
	"net"
	"strconv"
)

type TcpServer struct {
	conn_type string
	conn_host string
	conn_port int

	listener net.Listener
	clients  []*TcpClient
}

func CreateServer(ctype string, host string, port int) *TcpServer {
	return &TcpServer{
		conn_type: ctype,
		conn_host: host,
		conn_port: port,
		clients:   make([]*TcpClient, 0),
	}
}

func (server *TcpServer) Run() {
	fullhost := server.conn_host + ":" + strconv.Itoa(server.conn_port)
	listener, err := net.Listen(server.conn_type, fullhost)
	if err != nil {
		log.Fatal("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	server.listener = listener

	log.Println("Listening on ", fullhost)
	server.BeginAccept()
}

func (server *TcpServer) BeginAccept() {
	for {
		conn, err := server.listener.Accept()
		if err != nil {
			log.Println("Error accepting : ", err.Error())
			continue
		}

		log.Println("User accepted : ", conn.RemoteAddr().String())
		InitClient(conn)

		// append(server.clients, client)

		log.Printf("%d user(s) connected \n", len(server.clients))
	}
}
