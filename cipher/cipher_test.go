package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	IVval            = make([]byte, aes.BlockSize)
	cipherBlockErr   = errors.New("TEST : Couldn't Creat Cipher Block")
	ioReadFullErr    = errors.New("TEST : Failed to fill rand value")
	tempAESNewCipher = aesNewCipherVar
)

// func setTempIOReadFull() {
// 	 = func(n int, err error) {
// 		return -1, errors.New("IO Full")
// 	}
// }

func setAESNewCipherBlock() {
	aesNewCipherVar = func(key []byte) (cipher.Block, error) {
		return nil, cipherBlockErr
	}
}

var tst = []struct {
	key   string
	exErr error
}{
	{"key123", nil},
	{"ke123", cipherBlockErr},
}

func TestEncryptStream(t *testing.T) {
	for _, item := range tst {
		_, err := encryptStream(item.key, IVval)

		assert.Equalf(t, item.exErr, err, "Expected %v but got %v", item.exErr, err)
		defer func() {
			aesNewCipherVar = tempAESNewCipher
		}()
		setAESNewCipherBlock()
	}
}

func TestDecryptStream(t *testing.T) {
	for _, item := range tst {
		_, err := decryptStream(item.key, IVval)

		assert.Equalf(t, item.exErr, err, "Expected %v but got %v", item.exErr, err)
		defer func() {
			aesNewCipherVar = tempAESNewCipher
		}()

		setAESNewCipherBlock()
	}
}

func TestEncryptWriter(t *testing.T) {
	tempIOReadFull := ioReadFullVar
	f, err := os.OpenFile("test_secret_file", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		t.Errorf("TestEncryptWriter Failed to open Test File : %q", err)
	}

	_, err = EncryptWriter("key123", f)
	assert.Equalf(t, nil, err, "Expected %v but got %v", nil, err)

	ioReadFullVar = func(r io.Reader, buf []byte) (n int, err error) {
		return -1, ioReadFullErr
	}
	_, err = EncryptWriter("key123", f)
	assert.Equalf(t, ioReadFullErr, err, "Expected %q but got %b", ioReadFullErr, err)
	ioReadFullVar = tempIOReadFull

	setAESNewCipherBlock()
	_, err = EncryptWriter("key123", f)
	assert.Equalf(t, cipherBlockErr, err, "Expedcted %q but got %q", cipherBlockErr, err)
	aesNewCipherVar = tempAESNewCipher

	f.Close()

	_, err = EncryptWriter("key123", f)
	assert.Equalf(t, "encrypt: unable to write full iv to writer", err.Error(), "Expected %v but got %v", nil, err.Error())
}

func TestDecryptReader(t *testing.T) {

	f, err := os.OpenFile("test_secret_file", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		t.Errorf("Test Decrypt Reader failed to open file : %v ", err)
	}

	_, err = DecryptReader("key123", f)
	assert.Equalf(t, nil, err, "Expected %q but got %q", nil, err)

	_, err = DecryptReader("key123", f)
	assert.NotEqualf(t, nil, err, "Expected %q but got %q", nil, err)

	f.Close()

	f, err = os.OpenFile("test_secret_file", os.O_CREATE|os.O_RDWR, 0755)
	aesNewCipherVar = func(key []byte) (cipher.Block, error) {
		return nil, cipherBlockErr
	}
	_, err = DecryptReader("key123", f)
	assert.Equalf(t, cipherBlockErr, err, "Expected %q but got %q", cipherBlockErr, err)
}
