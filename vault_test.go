package secret

import (
	"crypto/cipher"
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mitchellh/go-homedir"
)

var (
	encodingKey       = "testkey_123"
	tempDecryptReader = DecryptReaderVar
	tempEncryptWriter = EncryptWriterVar
)

func setSecretPath(path string) string {
	dir, _ := homedir.Dir()
	return filepath.Join(dir, path)
}

func setEncryptWriter() {
	EncryptWriterVar = func(key string, w io.Writer) (*cipher.StreamWriter, error) {
		return nil, errors.New("Error While Encrypting")
	}
}

func setDecryptReader() {
	DecryptReaderVar = func(key string, r io.Reader) (*cipher.StreamReader, error) {
		return nil, errors.New("Error While Decrypting")
	}
}

func TestFile(t *testing.T) {
	path := setSecretPath("test_secrets")
	v := File(encodingKey, path)
	assert.Equal(t, encodingKey, v.encodingKey, "Expected %q but got %q", encodingKey, v.encodingKey)
}

func TestSave(t *testing.T) {
	filepath := setSecretPath("test_secrets")
	v := File(encodingKey, filepath)
	err := v.save()
	var str string
	assert.Equal(t, nil, err, "Expected %v got %v", str, err)

	setEncryptWriter()
	err = v.save()
	assert.Equalf(t, "Error While Encrypting", err.Error(), "Expected %v but got %v", "Error While Encrypting", err.Error())
	EncryptWriterVar = tempEncryptWriter

	val := File(encodingKey, "")
	err = val.save()
	assert.Equalf(t, "open : no such file or directory", err.Error(), "Expected %v but got %v", "open : no such file or directory", err.Error())
}

func TestLoad(t *testing.T) {
	filepath := setSecretPath("test_secrets")
	v := File(encodingKey, filepath)
	err := v.load()
	assert.Equalf(t, nil, err, "Expected %v but got %v", nil, err)

	setDecryptReader()

	err = v.load()

	assert.Equalf(t, "Error While Decrypting", err.Error(), "Expected %v but got %v", "Error While Decrypting", err.Error())

	DecryptReaderVar = tempDecryptReader

	val := File(encodingKey, "")

	err = val.load()

	assert.Equalf(t, nil, err, "Expected %v but got %v", nil, err)

}

func TestSet(t *testing.T) {
	filepath := setSecretPath("test_secrets")
	v := File(encodingKey, filepath)
	os.Remove(filepath)

	err := v.Set("test_demo_key", "Some_demo_value")
	assert.Equal(t, nil, err, "Expected %v but got %v", nil, err)

	setDecryptReader()
	err = v.Set("test_demo_key", "Some_demo_value")
	assert.Equalf(t, "Error While Decrypting", err.Error(), "Expected %v but got %v", "Error While Decrypting", err.Error())
	DecryptReaderVar = tempDecryptReader

	setEncryptWriter()
	err = v.Set("test_demo_key", "Some_demo_value")
	assert.Equalf(t, "Error While Encrypting", err.Error(), "Expected %v but got %v", "Error While Encrypting", err.Error())
	EncryptWriterVar = tempEncryptWriter
}

func TestGet(t *testing.T) {
	filepath := setSecretPath("test_secrets")
	v := File(encodingKey, filepath)

	value, _ := v.Get("test_demo_key")
	assert.Equalf(t, "Some_demo_value", value, "Expected %v but got %v", "Some_demo_value", value)

	setDecryptReader()
	_, err := v.Get("test_demo_key")
	assert.Equalf(t, "Error While Decrypting", err.Error(), "Expected %v but got %v", "Error While Decrypting", err.Error())
	DecryptReaderVar = tempDecryptReader

	value, err = v.Get("sdasda")
	assert.Equalf(t, "secret: no value for that key", err.Error(), "Expected %v but got %v", "secret: no value for that key", err.Error())
}
