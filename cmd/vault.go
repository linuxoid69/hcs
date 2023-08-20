/*
Copyright © 2023 Rustam Tagaev linuxoid69@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/linuxoid69/hcs/internal/helpers"
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
	Short: "Vault environment",
	Run: func(cmd *cobra.Command, args []string) {

		if addVault {
			addEntry()
			os.Exit(0)
		} else if deleteVault {
			helpers.DefaultHelp(cmd, &args)
			deleteEntry(args[0])
			os.Exit(0)
		} else if listVault {
			fmt.Println(getProfiles(ServiceVault, ListNames))
			os.Exit(0)
		}

		helpers.DefaultHelp(cmd, &args)
	},
}

func init() {
	rootCmd.AddCommand(vaultCmd)

	vaultCmd.PersistentFlags().BoolVar(&listVault, "list", false, "List Vault entries")
	vaultCmd.PersistentFlags().BoolVar(&addVault, "add", false, "Add new Vault entry")
	vaultCmd.PersistentFlags().BoolVar(&deleteVault, "delete", false, "Delete Vault entry (--delete foo)")
}

func getProfiles(service, key string) string {
	value, err := keychain.GetCredentials(service, key)
	if err != nil {
		fmt.Println("Profiles not exists yet")
		os.Exit(0)
	}

	return strings.ReplaceAll(strings.TrimLeft(value, " "), " ", "\n")
}

func deleteEntry(name string) {
	for _, i := range []string{"token", "host"} {
		keychain.DeleteCredentials(ServiceVault+"/"+name, i)
	}

	profileList, _ := keychain.GetCredentials(ServiceVault, ListNames)

	keychain.SetCredentials(
		&keychain.Secret{
			Service: ServiceVault,
			Key:     ListNames,
			Value:   strings.ReplaceAll(profileList, name, ""),
		},
	)
}

func addEntry() {
	name := prompt.GetInputName(prompt.PromptContent{
		ErrorMsg:    "Name is invalid or name is already exists",
		Label:       "Enter name for new entry:",
		ServiceName: ServiceVault,
		ServicePath: ListNames,
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

	profileList, _ := keychain.GetCredentials(ServiceVault, ListNames)

	keychain.SetCredentials(
		&keychain.Secret{
			Service: ServiceVault,
			Key:     ListNames,
			Value:   strings.Join(append(strings.Split(profileList, " "), name), " "),
		},
	)
}
