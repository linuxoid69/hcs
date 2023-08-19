/*
Copyright © 2023 Rustam Tagaev linuxoid69@gmail.com
*/
package helpers

import (
	"regexp"
	"strings"

	"github.com/linuxoid69/hcs/internal/keychain"
)

// IsValidHost - valid host or not.
func IsValidHost(host string) bool {
	matcher := regexp.MustCompile("^(http|https)://.+\\.+\\w+")

	return matcher.Match([]byte(host))
}

// IsValidName - valid name or not.
func IsValidName(name string) bool {
	if name == "" {
		return false
	} else {
		matcher := regexp.MustCompile("^.+\\s.*")

		if matcher.Match([]byte(name)) {
			return false
		}
	}

	return true
}

// IsNameExists - check name exists or not.
func IsNameExists(service, name string) bool {
	vaultList, _ := keychain.GetCredentials(service, name)

	for _, i := range strings.Split(vaultList, " ") {
		if i == name {
			return true
		}
	}

	return false
}
