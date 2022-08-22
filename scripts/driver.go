package main

import (
	// "fmt"
	"iq-bot/aws"
	"iq-bot/core"
	// "time"
)

func main() {

	cliEnv, err := core.LoadEnv()
	core.PrintIfErr(err)
	core.Success("environment : ", cliEnv)

	//connect to website
	connect, err := core.Login(core.WebsiteLogin{Url: cliEnv.Url, Username: cliEnv.Username, Password: cliEnv.Password})
	core.PrintIfErr(err)
	core.Success("connection : ", connect)
	
	aws.NavIq(connect)
	elems:=aws.Scrape(connect)
	core.Success("elem 1 : ", elems[0])

	// //visit website
	// connect, err := core.Visit("https://us-east-1.console.aws.amazon.com/")
	// core.PrintIfErr(err)
	// core.Success("connection : ", connect)
	

}