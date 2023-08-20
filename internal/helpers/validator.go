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
	matcher := regexp.MustCompile(`^(http|https):\/\/.+\.+\w+`)

	return matcher.Match([]byte(host))
}

// IsValidName - valid name or not.
func IsValidName(name string) bool {
	if name == "" {
		return false
	} else {
		matcher := regexp.MustCompile(`^.+\s.*`)

		if matcher.Match([]byte(name)) {
			return false
		}
	}

	return true
}

type Service struct {
	ServiceName string
	ServicePath string
	ProfileName string
}

// IsNameExists - check name exists or not.
func IsNameExists(service *Service) bool {
	profileList, _ := keychain.GetCredentials(service.ServiceName, service.ServicePath)

	for _, i := range strings.Split(profileList, " ") {
		if i == service.ProfileName {
			return true
		}
	}

	return false
}
