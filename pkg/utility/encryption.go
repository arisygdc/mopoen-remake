package utility

import (
	"crypto/sha1"
	"io"

	"golang.org/x/crypto/hkdf"
)

func HKDF16(secret, salt, info string) []byte {
	hashFunc := sha1.New
	kdf := hkdf.New(hashFunc, []byte(secret), []byte(salt), []byte(info))
	key := make([]byte, 16)
	_, _ = io.ReadFull(kdf, key)
	return key
}
