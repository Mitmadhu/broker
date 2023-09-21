package helper

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GetSalt() string{
	saltBytes := make([]byte, 16) 
	_, err := rand.Read(saltBytes)

	if err != nil {
		panic(fmt.Sprintf("error while generating salt, err: %v", err.Error()))
	}
	return hex.EncodeToString(saltBytes)
}