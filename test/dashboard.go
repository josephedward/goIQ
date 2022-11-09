package main

import (

	"iq-bot/core"
	"iq-bot/cli"
	"iq-bot/iq"
	"os"
	"reflect"
	"strconv"
	"time"
)

func main() {
	// fmt.Println("Arg length is %d", argLength)
	u := os.Args[2]

	browser := core.Manual(u)

	//navigate to IQ
	connect := core.Connect(browser, "https://iq.aws.amazon.com/work/#/requests")
	cli.Success("connection : ", connect)

	//takes a second, I guess
	time.Sleep(time.Second * 2)

	reqs := iq.GetRequests(connect)

	for _, req := range reqs {
		cli.Success("request : ", req)
		//get value of title and content
		title := reflect.ValueOf(req).FieldByName("title").String()
		content := reflect.ValueOf(req).FieldByName("content").String()
		author := reflect.ValueOf(req).FieldByName("author").String()
		cli.Success("title : ", title)
		cli.Success("author : ", author)
		cli.Success("content : ", content)
	}

	cli.Success("# of requests: ", strconv.Itoa(len(reqs)))
}
