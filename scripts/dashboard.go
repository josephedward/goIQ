package main

import (
	"iq-bot/core"
	"iq-bot/iq"
	"os"
	"reflect"
	"time"
)

func main() {
	// fmt.Println("Arg length is %d", argLength)
	u := os.Args[2]

	//navigate to IQ
	connect, err := core.Connect(u, "https://iq.aws.amazon.com/work/#/requests")
	core.PrintIfErr(err)
	core.Success("connection : ", connect)

	//takes a second, I guess
	time.Sleep(time.Second * 2)

	reqs := iq.GetRequests(connect)

	for _, req := range reqs {
		core.Success("request : ", req)
		//get value of title and content
		title := reflect.ValueOf(req).FieldByName("title").String()
		content := reflect.ValueOf(req).FieldByName("content").String()
		core.Success("title : ", title)
		core.Success("content : ", content)
	}
}
