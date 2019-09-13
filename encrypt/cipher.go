package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

// Encrypt will take in a String Key and a Plaintext and will convert it to a
// Hex represented value
func Encrypt(key, plaintext string) (string, error) {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	block, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.
	return fmt.Sprintf("%x", ciphertext), nil
}

// Decrypt will take in a Key and a hex representation of the ciphertext
// and Decrpyt it. Will return a string
// This code is Based of Examples from Go Doc
func Decrypt(key string, cipherHex string) (string, error) {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.

	block, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}

	//cipherHex :=
	//key, _ := hex.DecodeString("6368616e676520746869732070617373")
	ciphertext, _ := hex.DecodeString(cipherHex)

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("Cipher too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()

	fmt.Fprint(hasher, key)

	cipherKey := hasher.Sum(nil)
	block, err := aes.NewCipher(cipherKey)
	if err != nil {
		return nil, err
	}
	return block, nil
}
