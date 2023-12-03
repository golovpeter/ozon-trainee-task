package common

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateAlias(input string) string {
	const keyLength = 10

	hasher := sha256.New()

	hasher.Write([]byte(input))
	hashInBytes := hasher.Sum(nil)
	hash := hex.EncodeToString(hashInBytes)[:keyLength]

	return hash
}
