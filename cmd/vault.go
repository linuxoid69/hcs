/*
Copyright © 2023 Rustam Tagaev linuxoid69@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/linuxoid69/hcs/internal/keychain"
	"github.com/linuxoid69/hcs/internal/prompt"
	"github.com/spf13/cobra"
)

var listVault bool
var addVault bool
var deleteVault bool

const ServiceVault = "hcs/vault"
const ListNames = "list_names"

var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "HCS is an environment switch",
	Long:  `HCS is a utility for switching between different Vault environments`,
	Run: func(cmd *cobra.Command, args []string) {
		if addVault {
			addVaultEntry()
		} else if deleteVault {
			if len(args) == 0 {
				fmt.Println("Enter name, please!")
				os.Exit(1)
			}
			deleteEntryVault(args[0])
		} else if listVault {
			fmt.Println(getlistVault(ServiceVault, ListNames))
		}
	},
}

func init() {
	rootCmd.AddCommand(vaultCmd)

	vaultCmd.PersistentFlags().BoolVar(&listVault, "list", false, "list Vault entries")
	vaultCmd.PersistentFlags().BoolVar(&addVault, "add", false, "print 'add' to add new Vault entry")
	vaultCmd.PersistentFlags().BoolVar(&deleteVault, "delete", false, "delete Vault entry")
}

func getlistVault(service, key string) string {
	value, err := keychain.GetCredentials(service, key)
	if err != nil {
		fmt.Println("Can't get list Vault")
		os.Exit(1)
	}

	return strings.ReplaceAll(strings.TrimLeft(value, " "), " ", "\n")
}

func deleteEntryVault(name string) {
	for _, i := range []string{"token", "host"} {
		keychain.DeleteCredentials(ServiceVault+"/"+name, i)
	}

	vaultList, _ := keychain.GetCredentials(ServiceVault, ListNames)

	keychain.SetCredentials(
		&keychain.Secret{
			Service: ServiceVault,
			Key:     ListNames,
			Value:   strings.ReplaceAll(vaultList, name, ""),
		},
	)
}

func addVaultEntry() {
	name := prompt.GetInputName(prompt.PromptContent{
		ErrorMsg: "Name is invalid or name is already exists",
		Label:    "Enter name for new entry:",
		Service:  ServiceVault,
	})

	host := prompt.GetInputHost(prompt.PromptContent{
		ErrorMsg: "Host is invalid",
		Label:    "Enter Vault host e.g https://example.com:",
	})

	token := prompt.GetInputToken(prompt.PromptContent{
		ErrorMsg: "Token is invalid",
		Label:    "Enter token:",
	})

	keychain.SetCredentials(&keychain.Secret{
		Service: ServiceVault + "/" + name,
		Key:     "token",
		Value:   token,
	})

	keychain.SetCredentials(&keychain.Secret{
		Service: ServiceVault + "/" + name,
		Key:     "host",
		Value:   host,
	})

	vaultList, _ := keychain.GetCredentials(ServiceVault, ListNames)

	keychain.SetCredentials(
		&keychain.Secret{
			Service: ServiceVault,
			Key:     ListNames,
			Value:   vaultList + " " + name,
		},
	)
}
