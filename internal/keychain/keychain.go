/*
Copyright © 2023 Rustam Tagaev linuxoid69@gmail.com
*/
package keychain

import (
	"fmt"

	"github.com/zalando/go-keyring"
)

type Secret struct {
	Service string
	Key     string
	Value   string
}

// SetCredentials - set credentials.
func SetCredentials(secret *Secret) error {
	err := keyring.Set(secret.Service, secret.Key, secret.Value)
	if err != nil {
		return fmt.Errorf("Can't set secret: %w", err)
	}

	return nil
}

// GetCredentials - get credentials.
func GetCredentials(service string, key string) (string, error) {
	secret, err := keyring.Get(service, key)
	if err != nil {
		return "", fmt.Errorf("Can't get secret: %w", err)
	}

	return secret, nil
}

// DeleteCredentials - delete credentials.
func DeleteCredentials(service string, key string) error {
	if err := keyring.Delete(service, key); err != nil {
		return fmt.Errorf("Can't delete secret: %w", err)
	}

	return nil
}
