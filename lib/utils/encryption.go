package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func GenerateHash(text string) string {
	hash := sha256.New()
	hash.Write([]byte(text))
	str := hex.EncodeToString(hash.Sum(nil))
	return str
}

func Encrypt(data []byte, salt []byte) (string, error) {
	cipher_instance, err := aes.NewCipher(salt[:32])
	if err != nil {
		return "", err
	}

	gcm_instance, err := cipher.NewGCM(cipher_instance)
	if err != nil {
		return "", err
	}
	nonce_instance := make([]byte, gcm_instance.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce_instance)
	encrypted_data := gcm_instance.Seal(nonce_instance, nonce_instance, data, nil)
	return hex.EncodeToString(encrypted_data), nil
}

func Decrypt(data string, salt []byte) (string, error) {
	cipher_instance, err := aes.NewCipher(salt[:32])
	if err != nil {
		return "", err
	}

	gcm_instance, err := cipher.NewGCM(cipher_instance)
	if err != nil {
		return "", err
	}
	decoded, _ := hex.DecodeString(data)
	nonce, ciphered_text := decoded[:gcm_instance.NonceSize()], decoded[gcm_instance.NonceSize():]
	encrypted_data, err := gcm_instance.Open(nil, nonce, ciphered_text, nil)
	if err != nil {
		return "", err
	}
	return string(encrypted_data), nil
}
