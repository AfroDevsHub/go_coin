package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
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

func Encrypt(data []byte, salt string) (string, error) {
	byteInput := []byte(salt)
	md5Hash := md5.Sum(byteInput)
	cipher_instance, err := aes.NewCipher([]byte(hex.EncodeToString(md5Hash[:])))
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

func Decrypt(data []byte, salt string) (string, error) {
	val := md5.Sum([]byte(salt))
	cipher_instance, err := aes.NewCipher([]byte(hex.EncodeToString(val[:])))
	if err != nil {
		return "", err
	}

	gcm_instance, err := cipher.NewGCM(cipher_instance)
	if err != nil {
		return "", err
	}
	nonce, ciphered_text := data[:gcm_instance.NonceSize()], data[gcm_instance.NonceSize():]
	encrypted_data, err := gcm_instance.Open(nil, nonce, ciphered_text, nil)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(encrypted_data), nil
}
