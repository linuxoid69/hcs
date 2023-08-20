/*
Copyright © 2023 Rustam Tagaev linuxoid69@gmail.com
*/
package prompt

import (
	"errors"
	"fmt"
	"os"

	"github.com/linuxoid69/hcs/internal/helpers"
	"github.com/manifoldco/promptui"
)

const (
	prompt       = "{{ . }} "
	valid        = "{{ . | green }} "
	invalid      = "{{ . | red }} "
	success      = "{{ . | bold }} "
	mask         = '*'
	promptFailed = "Prompt failed %v\n"
)

type PromptContent struct {
	ErrorMsg    string
	Label       string
	ServiceName string
	ServicePath string
}

func GetInputName(pc PromptContent) string {
	validate := func(input string) error {
		if !helpers.IsValidName(input) {
			return errors.New(pc.ErrorMsg)
		}

		if helpers.IsNameExists(&helpers.Service{
			ServiceName: pc.ServiceName,
			ServicePath: pc.ServicePath,
			ProfileName: input,
		}) {
			return errors.New(pc.ErrorMsg)
		}

		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  prompt,
		Valid:   valid,
		Invalid: invalid,
		Success: success,
	}

	prompt := promptui.Prompt{
		Label:     pc.Label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf(promptFailed, err)
		os.Exit(1)
	}

	return result
}

func GetInputToken(pc PromptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.ErrorMsg)
		}

		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  prompt,
		Valid:   valid,
		Invalid: invalid,
		Success: success,
	}

	prompt := promptui.Prompt{
		Label:     pc.Label,
		Templates: templates,
		Validate:  validate,
		Mask:      mask,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf(promptFailed, err)
		os.Exit(1)
	}

	return result
}

func GetInputHost(pc PromptContent) string {
	validate := func(input string) error {
		if !helpers.IsValidHost(input) {
			return errors.New(pc.ErrorMsg)
		}

		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  prompt,
		Valid:   valid,
		Invalid: invalid,
		Success: success,
	}

	prompt := promptui.Prompt{
		Label:     pc.Label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf(promptFailed, err)
		os.Exit(1)
	}

	return result
}
