package crypt

import (
	"fmt"
	"crypto/rand"
	"encoding/hex"
)

// CreateHash return randoms 8 bytes arrays as a hex string 
func CreateHash() string {
	key := make([]byte,8);

	_, err := rand.Read(key);

	if err != nil {
		fmt.Println("Error when random keys: " ,err)
	}

	str := hex.EncodeToString(key)

	return str
}