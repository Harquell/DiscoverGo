package main

import (
	uuid "github.com/nu7hatch/gouuid"
	"lharquel.com/discovergo/utils"
)

type NetMessage struct {
	guid   uuid.UUID
	buffer []byte
}

func CreateNetMessage(buffer []byte) *NetMessage {
	return &NetMessage{
		guid:   utils.GenerateUUID(),
		buffer: buffer,
	}
}

func BufferToString(buffer []byte) string {
	return string(buffer)
}

func StringToBuffer(str string) []byte {
	return []byte(str)
}
