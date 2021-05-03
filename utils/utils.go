package utils

import uuid "github.com/nu7hatch/gouuid"

func GenerateUUID() uuid.UUID {
	uid, err := uuid.NewV4()
	if err != nil {
		return uuid.UUID{}
	}
	return *uid
}
