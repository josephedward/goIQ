package iq

import (
	"iq-bot/core"
	
// 	"github.com/go-rod/rod"
// 	// "github.com/go-rod/rod/lib/input"
// 	// "time"
// 	// "fmt"
)

type IqProvider struct {
	browserString string
	core.Connection
	requests []IqRequest
}

type IqRequest struct {
	title string
	content string
	author string
}

func GetRequests(connect core.Connection)(reqs []IqRequest){
	//scrape website
	elems:=GetTitles(connect)

	//get the content of each request
	for _, elem := range elems {
		content := GetContent(connect, elem)
		author := GetAuthor(connect, elem)
		reqs = append(reqs, IqRequest{title: elem.MustText(), content: content, author: author})
		InsertMessage(connect, elem, "Hello")
	}
	return reqs
}


func CreateProvider(browserString string)(provider IqProvider){
	//create a new provider
	provider = IqProvider{}
	provider.browserString = browserString
	provider.Connection = core.Connect(provider.browserString, "https://iq.aws.amazon.com/work/#/requests")
	return provider
}


// func (provider IqProvider) GetRequests()(reqs []IqRequest){
// 	return GetRequests(provider.Connection)
// }

// func SendMessage(connect core.Connection , elem *rod.Element, message string) {
// 	elem.MustClick()
// 	connect.Page.MustWaitLoad().MustElement("textarea[id^='initialResponse']").MustInput(message)
// 	core.Success("Message inserted")
// }