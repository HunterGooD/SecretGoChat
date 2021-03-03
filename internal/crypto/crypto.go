package crypto

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"io"

	"github.com/gofrs/uuid"
)

// GenerateName генерация имения для пользователей
func GenerateName() (string, error) {
	uuid, err := uuid.DefaultGenerator.NewV4()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

// GenerateBytes случайная полседовательность байтов
func GenerateBytes(max uint) []byte {
	var slice = make([]byte, max)
	if _, err := rand.Read(slice); err != nil {
		return nil
	}
	return slice
}

// GeneratePrivate приватный ключ для пользователя
func GeneratePrivate(bits uint) *rsa.PrivateKey {
	priv, err := rsa.GenerateKey(rand.Reader, int(bits))
	if err != nil {
		return nil
	}
	return priv
}

// HashPublic ...
func HashPublic(public *rsa.PublicKey) string {
	return Base64Encode(HashSum([]byte(StringPublic(public))))
}

// HashSum ...
func HashSum(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// ParsePublic ...
func ParsePublic(publicData string) *rsa.PublicKey {
	public, err := x509.ParsePKCS1PublicKey(Base64Decode(publicData))
	if err != nil {
		return nil
	}
	return public
}

// StringPublic ..
func StringPublic(public *rsa.PublicKey) string {
	return Base64Encode(x509.MarshalPKCS1PublicKey(public))
}

// EncryptRSA ..
func EncryptRSA(public *rsa.PublicKey, data []byte) []byte {
	result, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, public, data, nil)
	if err != nil {
		return nil
	}
	return result
}

// DecryptRSA ..
func DecryptRSA(priv *rsa.PrivateKey, data []byte) []byte {
	result, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, data, nil)
	if err != nil {
		return nil
	}
	return result
}

// Sign ...
func Sign(priv *rsa.PrivateKey, data []byte) []byte {
	sign, err := rsa.SignPSS(rand.Reader, priv, crypto.SHA256, data, nil)
	if err != nil {
		return nil
	}
	return sign
}

// Verify ...
func Verify(public *rsa.PublicKey, data, sign []byte) error {
	return rsa.VerifyPSS(public, crypto.SHA256, data, sign, nil)
}

// EnctyprAES ..
func EnctyprAES(key, data []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()
	data = paddingPKCS5(data, blockSize)
	cipherText := make([]byte, blockSize+len(data))
	iv := cipherText[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[blockSize:], data)
	return cipherText
}

// DecryptAES ...
func DecryptAES(key, data []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	blockSize := block.BlockSize()
	if len(data) < blockSize {
		return nil
	}
	iv := data[:blockSize]
	data = data[blockSize:]
	if len(data)%blockSize != 0 {
		return nil
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)
	return unpaddingPKS5(data)
}

func unpaddingPKS5(data []byte) []byte {
	length := len(data)
	if length == 0 {
		return nil
	}
	unpadding := int(data[length-1])
	if length < unpadding {
		return nil
	}
	return data[:(length - unpadding)]
}

func paddingPKCS5(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	paddText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, paddText...)
}

// Base64Decode ..
func Base64Decode(data string) []byte {
	result, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil
	}
	return result
}

// Base64Encode ..
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
