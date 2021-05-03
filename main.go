package main

const (
	CONN_HOST = "localhost"
	CONN_PORT = 8888
	CONN_TYPE = "tcp"
)

func main() {
	server := CreateServer(CONN_TYPE, CONN_HOST, CONN_PORT)

	server.Run()
}
