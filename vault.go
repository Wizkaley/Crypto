package secret

import (
	"secret/encrypt"
)

// Memory on Disk
func Memory(encodingKey string) Vault {
	return Vault{
		encodingKey: encodingKey,
		keyValues:   make(map[string]string),
	}
}

// Vault Struct
type Vault struct {
	encodingKey string
	keyValues   map[string]string
}

// Get a Vault value for the specified key
func (v *Vault) Get(key string) (string, error) {
	hex := v.keyValues[key]
	res, err := encrypt.Decrypt(key, hex)
	if err != nil {
		return "", err
	}

	return res, nil
}

// Set a Vault Value for with a Key
func (v *Vault) Set(key, text string) error {
	//v.keyValues[key] = text

	res, err := encrypt.Encrypt(key, text)
	if err != nil {
		return err
	}
	v.keyValues[key] = res
	return nil
}
