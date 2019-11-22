package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

// Encrypt will get an encrypted Reader
func Encrypt(key, id string, r io.Reader) (er io.Reader, err error) {
	// Get encrypt key via HMAC-SHA-256 based PBKDF2 algo.
	dk := pbkdf2.Key([]byte(key), []byte(id), 4096, 32, sha256.New)
	// Get IV via HMAC-SHA-1 based PBKDF2 algo.
	iv := pbkdf2.Key([]byte(key), []byte(id), 4096, aes.BlockSize, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		return nil, fmt.Errorf("encrypt: %w", err)
	}
	block.BlockSize()

	er = cipher.StreamReader{
		S: cipher.NewCFBEncrypter(block, iv),
		R: r,
	}
	return er, nil
}

// Decrypt will get an decrypted Reader
func Decrypt(key, id string, r io.Reader) (er io.Reader, err error) {
	// Get encrypt key via HMAC-SHA-256 based PBKDF2 algo.
	dk := pbkdf2.Key([]byte(key), []byte(id), 4096, 32, sha256.New)
	// Get IV via HMAC-SHA-1 based PBKDF2 algo.
	iv := pbkdf2.Key([]byte(key), []byte(id), 4096, aes.BlockSize, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		return nil, fmt.Errorf("decrypt: %w", err)
	}

	er = cipher.StreamReader{
		S: cipher.NewCFBDecrypter(block, iv),
		R: r,
	}
	return er, nil
}
