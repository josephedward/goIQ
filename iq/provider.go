package iq

import (
	"fmt"
	"iq-bot/core"
	"iq-bot/cli"
	"github.com/go-rod/rod"
)

// type IqProxy interface {
// 	CreateProvider(browserString string) IqProvider
// 	Login(connect core.Connection, login core.WebsiteLogin)
// 	GetRequests(connect core.Connection) []IqRequest //defaults to 20 requests
// 	GetBatchRequests(connect core.Connection, number int) []IqRequest
// }

type IqProvider struct {
	core.AwsEnv
	core.Connection
	Requests []IqRequest
}

type IqRequest struct {
	element *rod.Element
	title   string
	content string
	author  string
}


func GetRequests(connect core.Connection) (reqs []IqRequest) {
	//scrape website
	elems := GetTitles(connect)
	fmt.Println("elems : ", elems)

	//get the content of each request
	for _, elem := range elems {
		content := GetContent(connect, elem)
		author := GetAuthor(connect, elem)
		reqs = append(reqs, IqRequest{element: elem, title: elem.MustText(), content: content, author: author})
	}
	return reqs
}

func GetBatchRequests(connect core.Connection, number int) (reqs []IqRequest) {
	reqs = []IqRequest{}
	for {
		// do something
		if len(reqs) < number {
			reqs = append(reqs, GetRequests(connect)...)
		} else if len(reqs) == number {
			break
		} else {
			reqs = reqs[:number]
			break
		}
	}
	return reqs
}

func GetTitles(connect core.Connection) rod.Elements {
	fmt.Println("connect : ", connect)
	elems := connect.Page.MustWaitLoad().MustElements("div[class^='ProjectRailItem']")
	fmt.Println("elems : ", elems)
	return elems
}

func GetContent(connect core.Connection, elem *rod.Element) string {
	elem.MustClick()
	temp := connect.Page.MustWaitLoad().MustElement("div[class^='TextMessage__text']").MustText()
	return temp
}

func GetAuthor(connect core.Connection, elem *rod.Element) string {
	elem.MustClick()
	temp := connect.Page.MustWaitLoad().MustElement("span[class^='author']").MustText()
	return temp
}

func InsertMessage(connect core.Connection, elem *rod.Element, message string) {
	elem.MustClick()
	textarea := connect.Page.MustWaitLoad().MustElement("textarea[id^='initialResponse']")
	textarea.MustSelectAllText().MustInput("")
	textarea.MustInput(message)
	cli.Success("Message inserted")
}

func DraftIntro() {}

func AnswerAllRequests(connect core.Connection, reqs []IqRequest) {
	// for _, req := range reqs {

	// 	// AnswerRequest(connect, req)
	// }
}
