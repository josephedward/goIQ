package cli

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
)

type PromptContent struct {
	Label string
	errorMsg string
}

type PromptOptions struct {
	Label string
	Key   int64
}

func PromptGetInput(pc PromptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label: pc.Label,
		// Templates: templates,
		Validate: validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Input: %s\n", result)
	return result
}

func PromptDownload() bool {
	willDownload := PromptGetInput(
		PromptContent{Label: "Download the sandbox credentials file in plaintext? (yes/no)"})
	if willDownload == "yes" {
		fmt.Println("Downloading Sandbox Credentials...")
		return true
	} else if willDownload == "no" {
		fmt.Println("Not downloading...")
		return false
	} else {
		fmt.Println("Invalid Answer")
		PromptDownload()
	}
	return false
}

func PromptFileName() string {
	filename := PromptGetInput(PromptContent{Label: "What would you like to name the file?"})
	return filename
}

func PromptFilePath() string {
	filepath := PromptGetInput(PromptContent{Label: "Where would you like to save the file to?"})
	return filepath
}

func PromptConfig() bool {
	willAppend := PromptGetInput(
		PromptContent{Label: "Would you like to append the sandbox credentials file to your AWS config file? (yes/no)"})
	if willAppend == "yes" {
		fmt.Println("Appending Sandbox Credentials to AWS configs...")
		return true

	} else if willAppend == "no" {
		fmt.Println("Not Appending to AWS configs...")
		return false
	} else {
		fmt.Println("Invalid Answer")
		PromptConfig()
	}
	return false
}

func PromptRepoOwner() (owner string, err error) {
	owner = PromptGetInput(PromptContent{Label: "What is the name of the repository owner?"})
	if owner == "" {
		err = errors.New("please enter a valid repository owner")
		PromptRepoOwner()
	}
	return owner, err
}

func PromptRepoName() (repo string, err error) {
	repo = PromptGetInput(PromptContent{Label: "What is the name of the repository?"})
	if repo == "" {
		err = errors.New("please enter a valid repository name")
		PromptRepoName()
	}
	return repo, err
}
