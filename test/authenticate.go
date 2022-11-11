package main

import (
	"iq-bot/core"
	"iq-bot/cli"
	// "iq-bot/iq"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	// "os"
	// "time"
)


func main() {
	//declare error
	//load login information into memory
	AwsEnv, err := core.LoadEnv()
	cli.Success("environment : ", AwsEnv)
	cli.PrintIfErr(err)


	u := launcher.MustResolveURL("")
	browser := rod.New().ControlURL(u).MustConnect()

	// cli.Success("Connection before: ", Connection)
	Connection := core.Connect(browser, AwsEnv.Url)
	cli.Success("Connection after: ", Connection)

	// //login to aws
	Connection, err = core.SimpleLogin(Connection, core.WebsiteLogin{AwsEnv.Url, AwsEnv.Username, AwsEnv.Password})
	//wait for 2fa - this is a hack for now, need to remove
	cli.Success("Connection: ", Connection)
	cli.PrintIfErr(err)
	cli.Success("...waiting on 2FA... (return to browser)")

}