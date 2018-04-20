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
	_, err := Encrypt(key, params, false)
	if err != nil {
		t.Error(err)
	}
}

func TestEncryptAndDecrypt(t *testing.T) {
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
}

func TestEncryptVerbose(t *testing.T) {
	// should successfully create a token with verbose enabled
	_, err := Encrypt(key, params, true)
	if err != nil {
		t.Error(err)
	}
}

func TestEncryptNonAlphanumericKey(t *testing.T) {
	// should fail to create a token when the key is not alphanumeric
	_, err := Encrypt("_"+key, params, false)
	if err == nil {
		t.Errorf("expected encrypting a token with a non alphanumeric key to return an error")
	}
}

func TestEncryptEmptyKey(t *testing.T) {
	// should fail to create a token when the key is empty
	_, err := Encrypt("", params, false)
	if err == nil {
		t.Errorf("expected encrypting a token to fail when the key is empty")
	}
}

func TestEncryptEmptyParams(t *testing.T) {
	// should fail to create a token when params is empty
	_, err := Encrypt(key, "", false)
	if err == nil {
		t.Errorf("expected encrypting a token to fail when params is empty")
	}
}

func TestEncryptLargeParams(t *testing.T) {
	// should fail when the params are longer than 512 characters
	random, err := createRandomBytes(257)
	longParams := hex.EncodeToString(random)

	_, err = Encrypt(key, longParams, false)
	if err == nil {
		t.Errorf("Encrypt did not fail to create token when params is longer than 512 characters")
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
}

func TestDecryptVerbose(t *testing.T) {
	// encrypt the params
	token, err := Encrypt(key, params, false)
	if err != nil {
		t.Error(err)
	}

	// decrypt the token
	result, err := Decrypt(key, token, true)
	if err != nil {
		t.Error(err)
	}

	// compare params to decrypted token
	if params != result {
		t.Errorf("expected decrypted token to match params")
	}
}

func TestDecryptCustomerPortalToken(t *testing.T) {
	// decrypt token created from the customer portal
	result, err := Decrypt(key, "yfuuiWuy8LMiNR1Au3b9-LSNln-X5W-enqvNBlhlpwQspOoLlMX4fIecVLTQJTLMGET14FtLxmp8U6zaDSq5eD-gYMHz9V0", false)
	if err != nil {
		t.Error(err)
	}

	if params != result {
		t.Errorf("expected decrypted customer token to match params")
	}
}

func TestDecryptNonAlphanumericKey(t *testing.T) {
	// encrypt the params
	token, err := Encrypt(key, params, false)
	if err != nil {
		t.Error(err)
	}

	// should fail to decrypt a token when the key is not alphanumeric
	_, err = Decrypt("_"+key, token, false)
	if err == nil {
		t.Errorf("expected decrypting a token with a non alphanumeric key to return an error")
	}
}

func TestDecryptEmptyKey(t *testing.T) {
	// encrypt the params
	token, err := Encrypt(key, params, false)
	if err != nil {
		t.Error(err)
	}

	// should fail to create a token when the key is empty
	_, err = Decrypt("", token, false)
	if err == nil {
		t.Errorf("expected decrypting a token to with an empty key to return an error")
	}
}

func TestDecryptGarbageToken(t *testing.T) {
	// should fail to create a token when the key is empty
	_, err := Decrypt(key, "yfuuiWuy8LMiNR1Au3b9-LSNln-X5W-", false)
	if err == nil {
		t.Errorf("expected decrypting a token to with an empty key to return an error")
	}
}
