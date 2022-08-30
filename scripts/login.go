package main 

import (
	"iq-bot/core"
	// "iq-bot/iq"
	// "fmt"
	"os"
	// "time"
)

func main(){
    // fmt.Println("Arg length is %d", argLength)
	u := os.Args[3]

	//load login information into memory
	cliEnv, err := core.LoadEnv()
	core.PrintIfErr(err)
	core.Success("environment : ", cliEnv)

	//connect to aws
	connect, err := core.Connect(u, cliEnv.Url)
	core.PrintIfErr(err)
	core.Success("connection : ", connect)

	//enter login credentials
	core.Login(connect, core.WebsiteLogin{cliEnv.Url, cliEnv.Username, cliEnv.Password})
	//(dont forget to manually enter 2fa)
}

