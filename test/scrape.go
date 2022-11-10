package main 

import (
	"iq-bot/core"
	"iq-bot/cli"
	"iq-bot/iq"
	"os"
	"time"
)

func main() {
	u := os.Args[2]
	browser := core.Manual(u)
	connect := core.Connect(browser, "https://iq.aws.amazon.com/work/#/requests")
	cli.Success("connection : ", connect)
	time.Sleep(time.Second * 2)
	elems := iq.GetElements(connect)
	iq.DisplayRequests(elems)
}