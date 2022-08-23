package main

import (
	// "fmt"
	// "iq-bot/iq"
	"iq-bot/core"
	// "time"
)

func main() {

	//load login information into memory
	cliEnv, err := core.LoadEnv()
	core.PrintIfErr(err)
	core.Success("environment : ", cliEnv)

	//connect to website
	connect, err := core.Login(core.WebsiteLogin{Url: cliEnv.Url, Username: cliEnv.Username, Password: cliEnv.Password})
	core.PrintIfErr(err)
	core.Success("connection : ", connect)
	
	// aws.NavIq(connect)
	// elems:=aws.Scrape(connect)
	// core.Success("elem 1 : ", elems[0])

	

}