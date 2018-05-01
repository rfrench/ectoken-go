package v3

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
)

const (
	ivSize      = 12
	authTagSize = 16
)

// Encrypt generates an version 3 encrypted authentication token
func Encrypt(key string, params string, verbose bool) (string, error) {
	// validate key
	if (len(key) <= 0) || (!isAlphanumeric(key)) {
		return "", errors.New("key must be at least 1 character in length and only contain alphanumeric characters")
	}

	// validate params length
	paramsLength := len(params)
	if (paramsLength <= 0) || (paramsLength > 512) {
		return "", errors.New("params must between 1 and 512 characters in length")
	}

	// sha256 hash key
	hash := hashKey(key)

	// create cipher
	block, err := aes.NewCipher(hash)
	if err != nil {
		return "", err
	}

	// create gcm
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// create iv
	iv, err := randomBytes(ivSize)
	if err != nil {
		return "", err
	}

	// encrypted token in bytes (iv + cipherText + tag)
	token := gcm.Seal(iv, iv, []byte(params), nil)

	// parse cipher text from token
	cipherText := token[ivSize:(len(token) - authTagSize)]

	// parse authorization tag from token
	tag := token[len(cipherText)+ivSize:]

	// verbose output
	if verbose == true {
		fmt.Printf("+-------------------------------------------------------------\n")
		fmt.Printf("| iv:                %s\n", hex.EncodeToString(iv))
		fmt.Printf("| cipherText:        %s\n", hex.EncodeToString(cipherText))
		fmt.Printf("| tag:               %s\n", hex.EncodeToString(tag))
		fmt.Printf("+-------------------------------------------------------------\n")
		fmt.Printf("| token:             %s\n", hex.EncodeToString(token))
		fmt.Printf("+-------------------------------------------------------------\n")
	}

	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(token), nil
}

// Decrypt decrypts a version 3 authentication token
func Decrypt(key string, token string, verbose bool) (string, error) {
	// validate key
	if (len(key) <= 0) || (!isAlphanumeric(key)) {
		return "", errors.New("key must be at least 1 character in length and only contain alphanumeric characters")
	}

	// sha256 hash key
	hash := hashKey(key)

	// create cipher
	block, err := aes.NewCipher(hash)
	if err != nil {
		return "", err
	}

	// create gcm
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// decode token
	decodedToken, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(token)
	if err != nil {
		return "", err
	}

	// parse iv
	iv := decodedToken[0:ivSize]

	// parse cipher text
	cipherText := decodedToken[ivSize:]

	// parse authorization tag from token
	tag := decodedToken[len(decodedToken)-authTagSize:]

	// decrypt params
	params, err := gcm.Open(nil, iv, cipherText, nil)
	if err != nil {
		return "", err
	}

	// verbose output
	if verbose == true {
		fmt.Printf("+-------------------------------------------------------------\n")
		fmt.Printf("| iv:                %s\n", hex.EncodeToString(iv))
		fmt.Printf("| cipherText:        %s\n", hex.EncodeToString(cipherText))
		fmt.Printf("| tag:               %s\n", hex.EncodeToString(tag))
		fmt.Printf("+-------------------------------------------------------------\n")
		fmt.Printf("| params:            %s\n", params)
		fmt.Printf("+-------------------------------------------------------------\n")
	}

	return string(params), nil
}

// random bytes used for initialization vectors
func randomBytes(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// SHA256 hashes the key
func hashKey(key string) []byte {
	h := sha256.New()
	h.Write([]byte(key))

	return h.Sum(nil)
}

// determines if a string is alphanumeric
func isAlphanumeric(s string) bool {
	r := regexp.MustCompile(`^[a-zA-Z0-9]*$`)
	return r.MatchString(s)
}
