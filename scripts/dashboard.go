package main 

import (
	"iq-bot/core"
	"iq-bot/iq"
	// "fmt"
	"os"
)

func main(){
    // fmt.Println("Arg length is %d", argLength)
	u := os.Args[3]

	// //load login information into memory
	// cliEnv, err := core.LoadEnv()
	// core.PrintIfErr(err)
	// core.Success("environment : ", cliEnv)

	// //connect to aws 
	// connect, err := core.Connect(u, cliEnv.Url)
	// core.PrintIfErr(err)
	// core.Success("connection : ", connect)
	
	// core.Login(connect, core.WebsiteLogin{cliEnv.Url, cliEnv.Username, cliEnv.Password})
		
	//navigate to IQ
	connect, err := core.Connect(u,"https://iq.aws.amazon.com/work/#/requests")
	core.PrintIfErr(err)
	core.Success("connection : ", connect)

	//scrape website
	elems:=iq.GetTitles(connect)
	
	//print out the titles
	for _, elem := range elems {
		core.Success("Title: ", elem.MustText())
	}
	
}

