package storage

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"time"

	"gopkg.in/night-codes/types.v1"

	"golang.org/x/crypto/bcrypt"
)

// GenerateToken returns a unique token based on the provided email string
func generateToken() string {
	rand.Seed(time.Now().UnixNano())
	hash, err := bcrypt.GenerateFromPassword([]byte(types.String(rand.Int())), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash to store:", string(hash))

	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}
