package main

import (
	"log"
	"net"

	uuid "github.com/nu7hatch/gouuid"
	"lharquel.com/discovergo/utils"
)

type TcpClient struct {
	uid  uuid.UUID
	name string
	conn net.Conn
}

func InitClient(conn net.Conn) *TcpClient {
	client := &TcpClient{
		uid:  utils.GenerateUUID(),
		name: "",
		conn: conn,
	}

	go client.ClientRoutine()
	return client
}

func (client *TcpClient) ClientRoutine() {
	defer client.Close()
	for {
		msg, err := client.ReadMessage()
		if err != nil {
			log.Println("User disconnected : ", client.conn.RemoteAddr().String())
			return
		}

		log.Printf("from %s[%s] > %s", client.name, client.conn.RemoteAddr().String(), BufferToString(msg.buffer))
	}
}

func (client *TcpClient) Close() {
	client.conn.Close()
}

func (client *TcpClient) ReadMessage() (*NetMessage, error) {
	buf := make([]byte, 1024)

	reqLen, err := client.conn.Read(buf)
	if err != nil {
		return &NetMessage{}, err
	}
	msg := CreateNetMessage(buf[0:reqLen])

	return msg, nil
}
