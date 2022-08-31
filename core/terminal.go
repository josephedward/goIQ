package core

import (
	"fmt"
	// "github.com/manifoldco/promptui"
)


var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func PrintIfErr(err error) {
	if err != nil {
		fmt.Println(Red + err.Error() + Reset)
	}
}

func Success(message ...interface{}) {
	fmt.Print(Green + message[0].(string) + Reset + " ")
	for _, msg := range message {
		//if msg is the first, don't print a newline
		if msg == message[0] {
			continue
		}
		fmt.Println(msg)
	}
}
