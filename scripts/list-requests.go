package main 

import (
	"iq-bot/core"
	"iq-bot/iq"
	// "fmt"
	"os"
)

func main(){
	// argLength := len(os.Args[1:])
        // use fmt.Printf() to format string
    // fmt.Println("Arg length is %d", argLength)
	u := os.Args[3]


	// 	// //visit website
	connect, err := core.VisitManual(u,"https://iq.aws.amazon.com/work/#/requests")
	core.PrintIfErr(err)
	core.Success("connection : ", connect)

	// // //scrape website
	elems:=iq.GetTitles(connect)
	
	//print out the titles
	for _, elem := range elems {
		core.Success("Title : ", elem.MustText())
	}
	
}

