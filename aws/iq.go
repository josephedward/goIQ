package aws

import (
	"iq-bot/core"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"time"
	"fmt"
)

// "div[class^='CopyableInstanceField__Value']")

func NavIq(connect core.Connection) {

	// globalNav-search
	elem := 	connect.Page.MustWaitLoad().MustElementR("input", "Search").MustInput("iq").MustType(input.Enter)
	core.Success("elem : ", elem)

	
	
	// elem = connect.Page.MustWaitLoad().MustElementR("a","dashboard").MustClick()
	// core.Success("elem : ", elem)
	// MustElement("a[href='/work']").MustClick()
	
	
}


func Scrape(connect core.Connection) rod.Elements {
	connect.Page.MustNavigate("https://iq.aws.amazon.com/work/#/requests")
	time.Sleep(5 * time.Second)
	elems := connect.Page.MustWaitLoad().MustElements("div[class^='ProjectRailItem__title']")
	fmt.Println(elems[0].MustText())
	return elems
}

