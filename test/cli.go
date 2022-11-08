package main

import (
	"errors"
	"fmt"
	"iq-bot/core"
	"iq-bot/iq"
	"os"
	"reflect"
	"strconv"
	"strings"
	"github.com/manifoldco/promptui"
)

type promptOptions struct {
	Label string
	Key   int64
}

type promptContent struct {
	Label string
	// Items    []string
	errorMsg string
}

func main() {
	var p iq.IqProvider
	// var p acloud.ACloudProvider
	Execute(p)
}

func Execute(p iq.IqProvider) {
	// fmt.Println("Arg length is %d", argLength)
	// u := os.Args[1]

	options := []promptOptions{
		{
			Label: "Exit CLI",
			Key:   0,
		},
		{
			Label: "Login to AWS",
			Key:   1,
		},
		{
			Label: "Get Requests",
			Key:   2,
		},
		{
			Label: "Send Message",
			Key:   3,
		},
		{
			Label: "Send Bulk Message",
			Key:   4,
		},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Label | cyan }} ",
		Inactive: "  {{ .Label | cyan }} ",
		Selected: "\U0001F336 {{ .Label | red | cyan }}",
	}

	searcher := func(input string, index int) bool {
		option := options[index]
		name := strings.Replace(strings.ToLower(option.Label), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Welcome to IQ-Bot. Please choose an option: ",
		Items:     options,
		Templates: templates,
		// Size:      4,
		Searcher: searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
		// return ACloudEnv{}, err
	}

	fmt.Printf("You choose number %d: %s\n", i+1, options[i].Label)

	switch options[i].Key {
	case 0:
		os.Exit(0)
	case 1:
		// ConnectIqProvider(p)
		LoginToAws(p)
	case 2:

	case 3:
		// reqs := GetBatchRequestNumber(p)
		// PrintRequests(reqs)

	}

	Execute(p)
}

func LoginToAws(p iq.IqProvider) (err error) {
	//load login information into memory
	p.AwsEnv, err = core.LoadEnv()
	core.PrintIfErr(err)
	cli.Success("environment : ", p.AwsEnv)

	browserino := core.CustomLaunch()
	p.Connection = core.Connect(browserino, p.AwsEnv.Url)

	//login to aws
	p.Login(p.Connection, p.AwsEnv)

	return err
}

func promptGetInput(pc promptContent) string {
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

func GetBatchRequestNumber(connect core.Connection) (reqs []iq.IqRequest) {
	numberRequests := promptGetInput(promptContent{
		Label: "Enter the number of requests you want to scrape",
	})
	num, _ := strconv.Atoi(numberRequests)
	reqs = iq.GetBatchRequests(connect, num)
	return reqs
}

func PrintRequests(reqs []iq.IqRequest) {
	for _, req := range reqs {
		cli.Success("Request:", reflect.ValueOf(req).FieldByName("title").String())
	}
}

func SendMessage(connect core.Connection, templates *promptui.SelectTemplates, searcher func(input string, index int) bool, reqs []iq.IqRequest) {

	// options := []iq.IqRequest{reqs...}
	options := []promptOptions{}
	for _, req := range reqs {
		cli.Success("Request:", reflect.ValueOf(req).FieldByName("title").String())
		// options = append(options, promptOptions{
		// 	Label: req.title,
		// 	Key:   int64(req.id),
		// })
	}

	prompt := promptui.Select{
		Label:     "Welcome to IQ-Bot. Please choose an option: ",
		Items:     options,
		Templates: templates,
		// Size:      4,
		Searcher: searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		main()
		// return ACloudEnv{}, err
	}

	fmt.Printf("You choose request %d: %s\n", i+1, options[i].Label)

	// iq.InsertMessage(connect, req.element, msg)

}

// func SendBulkMessage(connect core.Connection, reqs []iq.IqRequest) {
// 	for _, req := range reqs {
// 		iq.SendBulkMessage(connect, req)
// 	}
// }

// cli.Success("# of requests: ",strconv.Itoa(len(reqs)) )

// // func promptEnvFile() (ACloudEnv, error) {
// // 	//load env variables
// // 	env_path := promptGetInput(promptContent{
// // 		Label: "Please enter the path to the .env file from this directory",
// // 	})

// // 	env, err := getEnv(env_path)

// // 	return env, err
// // }

// // func promptManual() (ACloudEnv, error) {

// // 	tempEnv := ACloudEnv{}

// // 	// get env vars via cli prompt
// // 	tempEnv.Url = promptGetInput(promptContent{
// // 		Label: "Name of web property URL you would like to login to",
// // 	})
// // 	// if tempEnv.Url == "" {
// // 	// 	tempEnv.Url = "https://learn.acloud.guru/cloud-playground/cloud-sandboxes"
// // 	// }
// // 	tempEnv.Username = promptGetInput(promptContent{
// // 		Label: "What is your username",
// // 	})
// // 	tempEnv.Password = promptGetInput(promptContent{
// // 		Label: "What is your password",
// // 	})
// // 	// get aws_path via cli prompt
// // 	tempEnv.Aws_path = promptGetInput(promptContent{
// // 		Label: "Where would you like your sandbox credentials appended",
// // 	})
// // 	// get download path via cli prompt
// // 	tempEnv.Download_key = promptGetInput(promptContent{
// // 		Label: "What would you like the name of your sandbox credentials file to be",
// // 	})
// // 	//if all env vars are set, return the env
// // 	if tempEnv.Url != "" && tempEnv.Username != "" && tempEnv.Password != "" && tempEnv.Aws_path != "" && tempEnv.Download_key != "" {
// // 		return tempEnv, nil
// // 	} else {
// // 		fmt.Println("Please fill out all fields")
// // 		promptManual()
// // 	}
// // 	return tempEnv, nil
// // }

// // func PromptDownload() bool{
// // 	willDownload := promptGetInput(
// // 		promptContent{Label: "Would you like to download the sandbox credentials file in plaintext? (yes/no)"})
// // 	if willDownload == "yes" {
// // 		fmt.Println("Downloading Sandbox Credentials...")
// // 		return true
// // 	} else if willDownload == "no" {
// // 		fmt.Println("Not downloading...")
// // 		return false
// // 	} else {
// // 		fmt.Println("Invalid Answer")
// // 		PromptDownload()
// // 	}
// // 	return false
// // }

// // func PromptFileName() string {
// // 	filename := promptGetInput(promptContent{Label: "What would you like to name the file?"})
// // 	return filename
// // }

// // func PromptFilePath() string {
// // 	filepath := promptGetInput(promptContent{Label: "Where would you like to save the file to?"})
// // 	return filepath
// // }

// // func PromptConfig() bool{
// // 	willAppend := promptGetInput(
// // 		promptContent{Label: "Would you like to append the sandbox credentials file to your AWS config file? (yes/no)"})
// // 	if willAppend == "yes" {
// // 		fmt.Println("Appending Sandbox Credentials to AWS configs...")
// // 		return true

// // 	} else if willAppend == "no" {
// // 		fmt.Println("Not Appending to AWS configs...")
// // 		return false
// // 	} else {
// // 		fmt.Println("Invalid Answer")
// // 		PromptConfig()
// // 	}
// // 	return false
// // }
