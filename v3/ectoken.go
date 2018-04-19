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
	"strings"
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
	hash := createHash(key)

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
	iv, err := createRandomBytes(ivSize)
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

	return encode(token), nil
}

// Decrypt decrypts a version 3 authentication token
func Decrypt(key string, token string, verbose bool) (string, error) {
	// validate key
	if (len(key) <= 0) || (!isAlphanumeric(key)) {
		return "", errors.New("key must be at least 1 character in length and only contain alphanumeric characters")
	}

	// sha256 hash key
	hash := createHash(key)

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
	decodedToken, err := decode(token)
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

<<<<<<< HEAD
// base64 encodes and replaces characters just like urlsafe_b64encode
=======
>>>>>>> f7fcf6cf4894c1c5001e0911b00502c17657000a
func encode(b []byte) string {
	token := base64.StdEncoding.EncodeToString(b)
	token = strings.Replace(token, "=", "", -1)
	token = strings.Replace(token, "+", "-", -1)
	token = strings.Replace(token, "/", "_", -1)

	return token
}

<<<<<<< HEAD
// \re-replaces characters just like urlsafe_b64decode and base64 decodes string
=======
>>>>>>> f7fcf6cf4894c1c5001e0911b00502c17657000a
func decode(token string) ([]byte, error) {
	token = strings.Replace(token, "-", "+", -1)
	token = strings.Replace(token, "_", "/", -1)

	switch len(token) % 4 {
	case 2:
		token += "=="
	case 3:
		token += "="
	}

	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	return decoded, nil
}

<<<<<<< HEAD
// random bytes used for initialization vectors
=======
>>>>>>> f7fcf6cf4894c1c5001e0911b00502c17657000a
func createRandomBytes(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

<<<<<<< HEAD
// SHA256 hashes a string
=======
>>>>>>> f7fcf6cf4894c1c5001e0911b00502c17657000a
func createHash(s string) []byte {
	h := sha256.New()
	h.Write([]byte(s))

	return h.Sum(nil)
}

<<<<<<< HEAD
// determines if a string is alphanumeric
=======
>>>>>>> f7fcf6cf4894c1c5001e0911b00502c17657000a
func isAlphanumeric(s string) bool {
	r := regexp.MustCompile(`^[a-zA-Z0-9]*$`)
	return r.MatchString(s)
}
