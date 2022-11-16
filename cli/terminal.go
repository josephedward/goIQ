package cli

import (
	"fmt"
	"github.com/rs/zerolog"
	// "github.com/rs/zerolog/log"
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
		Error(err.Error())
	}
}

func Success(message ...interface{}) {

	//if log level is debug, print success messages
	if zerolog.GlobalLevel() == zerolog.DebugLevel {
		for _, msg := range message {
			s, ok := msg.(string) // the "ok" boolean will flag success.
			if ok {
				fmt.Println(Green + string(s) + Reset)
			} else {
				fmt.Println(msg)
			}
		}
	}
}

func Error(message ...interface{}) {
	//if log level is debug, print err messages
	if zerolog.GlobalLevel() == zerolog.DebugLevel {

		for _, msg := range message {
			s, ok := msg.(string) // the "ok" boolean will flag success.
			if ok {
				fmt.Println(Red + string(s) + Reset)
			} else {
				fmt.Println(msg)
			}
		}
	}
}

func Welcome() {
	fmt.Println(Cyan + `		 -------------------------------------------			` + Reset)
	fmt.Println(Green + `		 . d8888b.   .d88888b.  8888888  .d88888b.  			` + Reset)
	fmt.Println(Green + `		 d88P  Y88b d88P" "Y88b   888   d88P" "Y88b 			` + Reset)
	fmt.Println(Green + `		 888    888 888     888   888   888     888 			` + Reset)
	fmt.Println(Green + `		 888        888     888   888   888     888 			` + Reset)
	fmt.Println(Green + `		 888  88888 888     888   888   888     888 			` + Reset)
	fmt.Println(Green + `		 888    888 888     888   888   888 Y8b 888 			` + Reset)
	fmt.Println(Green + `		 Y88b  d88P Y88b. .d88P   888   Y88b.Y8b88P 			` + Reset)
	fmt.Println(Green + `		  "Y8888P88  "Y88888P"  8888888  "Y888888"  			` + Reset)
	fmt.Println(Green + `		                                       Y8b  			` + Reset)
	fmt.Println(Cyan + `		 -------------------------------------------			` + Reset)
	fmt.Println(Yellow + `		 The AWS IQ Bot for Conversion and Retention			` + Reset)
	fmt.Println(Cyan + `		 -------------------------------------------			` + Reset)
}
