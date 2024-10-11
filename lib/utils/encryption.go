package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
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
	nonce := salt[:gcm_instance.NonceSize()]
	encrypted_data := gcm_instance.Seal(nonce, nonce, data, nil)
	// combined := append(nonce, encrypted_data...)
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
	ciphered_text := decoded[gcm_instance.NonceSize():]
	encrypted_data, err := gcm_instance.Open(nil, decoded[:gcm_instance.NonceSize()], ciphered_text, nil)
	if err != nil {
		return "", err
	}
	return string(encrypted_data), nil
}
