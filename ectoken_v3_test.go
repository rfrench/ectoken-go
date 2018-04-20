package v3

import (
	"encoding/hex"
	"testing"
)

const (
	key    = "GF8PHCp3xy5ypSaJKmPMH2M4"
	params = "ec_expire=1257642471&ec_clientip=11.22.33.1"
)

func TestEncrypt(t *testing.T) {
	// encrypt the params
	token, err := Encrypt(key, params, false)
	if err != nil {
		t.Error(err)
	}

	// decrypt the token
	result, err := Decrypt(key, token, false)
	if err != nil {
		t.Error(err)
	}

	// compare params to decrypted token
	if params != result {
		t.Errorf("expected decrypted token to match params")
	}

	// should fail to create a token when the key is not alphanumeric
	token, err = Encrypt("_"+key, params, false)
	if err == nil {
		t.Errorf("expected encrypting a token with a non alphanumeric key to return an error")
	}

	// should fail to create a token when the key is empty
	token, err = Encrypt("", params, false)
	if err == nil {
		t.Errorf("expected encrypting a token to fail when the key is empty")
	}

	// should fail when the params are longer than 512 characters
	random, err := createRandomBytes(257)
	longParams := hex.EncodeToString(random)

	token, err = Encrypt(key, longParams, false)
	if err == nil {
		t.Errorf("Encrypt did not fail to create token when params is longer than 512 characters")
	}

	token, err = Encrypt(key, "", false)
	if err == nil {
		t.Errorf("Encrypt did not fail to create token when params is empty")
	}
}

func TestDecrypt(t *testing.T) {
	// encrypt the params
	token, err := Encrypt(key, params, false)
	if err != nil {
		t.Error(err)
	}

	// decrypt the token
	result, err := Decrypt(key, token, false)
	if err != nil {
		t.Error(err)
	}

	// compare params to decrypted token
	if params != result {
		t.Errorf("expected decrypted token to match params")
	}

	// decrypt token created from the customer portal
	result, err = Decrypt(key, "yfuuiWuy8LMiNR1Au3b9-LSNln-X5W-enqvNBlhlpwQspOoLlMX4fIecVLTQJTLMGET14FtLxmp8U6zaDSq5eD-gYMHz9V0", false)
	if params != result {
		t.Errorf("expected decrypted customer token to match params")
	}

	// should fail to decrypt a token when the key is not alphanumeric
	result, err = Decrypt("_"+key, params, false)
	if err == nil {
		t.Errorf("expected decrypting a token with a non alphanumeric key to return an error")
	}

	// should fail to create a token when the key is empty
	result, err = Decrypt("", params, false)
	if err == nil {
		t.Errorf("expected decrypting a token to with an empty key to return an error")
	}
}
