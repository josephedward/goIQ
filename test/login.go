package main

import (
	"iq-bot/core"
	"iq-bot/cli"
	// "iq-bot/iq"
	// "fmt"
	"os"
	// "time"
)

func main() {
	// fmt.Println("Arg length is %d", argLength)
	u := os.Args[3]

	//load login information into memory
	cliEnv, err := core.LoadEnv()
	cli.PrintIfErr(err)
	cli.Success("environment : ", cliEnv)

	browser := core.Manual(u)

	//connect to aws
	connect := core.Connect(browser, cliEnv.Url)
	cli.PrintIfErr(err)
	cli.Success("connection : ", connect)

	//enter login credentials
	core.SimpleLogin(connect, core.WebsiteLogin{cliEnv.Url, cliEnv.Username, cliEnv.Password})
	//(dont forget to manually enter 2fa)
}
