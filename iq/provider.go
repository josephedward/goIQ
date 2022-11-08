package iq

import (
	"fmt"
	"github.com/go-rod/rod"
	"iq-bot/cli"
	"iq-bot/core"
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
	budget  string
	label   string
	header  string
}

// func GetRequests(connect core.Connection) (reqs []IqRequest) {
// 	//scrape website
// 	elems := GetTitles(connect)
// 	fmt.Println("elems : ", elems)

// 	//get the content of each request
// 	for _, elem := range elems {
// 		content := GetContent(connect, elem)
// 		author := GetAuthor(connect, elem)

// 	}
// 	return reqs
// }

func GetBatchRequests(connect core.Connection, number int) (reqs []IqRequest) {
	reqs = []IqRequest{}
	for {
		// do something
		if len(reqs) < number {
			reqs = append(reqs, GetElements(connect)...)
		} else if len(reqs) == number {
			break
		} else {
			reqs = reqs[:number]
			break
		}
	}
	return reqs
}

func GetElements(connect core.Connection) (reqs []IqRequest) {
	// cli.Success("connect : ", connect)

	//get whole element
	elems := connect.Page.MustWaitLoad().MustElements("div[class^='content ProjectRailItem']")
	// cli.Success("elems : ", elems)

	headers := connect.Page.MustWaitLoad().MustElements("div[class^='ProjectRailItem__header']")
	titles := connect.Page.MustWaitLoad().MustElements("div[class^='ProjectRailItem__title']")
	labels := connect.Page.MustWaitLoad().MustElements("div[class^='ProjectRailItem__label']")
	budgets := connect.Page.MustWaitLoad().MustElements("div[class^='ProjectRailItem__budget']")

	//loop over length of elems
	for i := 0; i < len(elems); i++ {
		cli.Success("title : ", titles[i].MustText())
		cli.Success("header : ", headers[i].MustText())
		cli.Success("label : ", labels[i].MustText())
		cli.Success("budget : ", budgets[i].MustText())
		reqs = append(reqs, IqRequest{
			element: elems[i],
			title:  titles[i].MustText(),
			header: headers[i].MustText(),
			
			// use conditional for label 
			label: labels[i].MustText(),

			//use condtional for budget
			budget: budgets[i].MustText()})
	}

	//for each elem, get the text of the header, title, label, and budget
	//create a new IqRequest object with the text of each element
	//append the new IqRequest object to the array of IqRequest objects
	//return the array of IqRequest objects

	return reqs
}

func GetRequestDetails(connect core.Connection, req IqRequest) {
	//get the title, content, author, budget, label, and header
	req.element.MustClick()

}

func GetTitles(connect core.Connection) rod.Elements {
	fmt.Println("connect : ", connect)
	elems := connect.Page.MustWaitLoad().MustElements("div[class^='ProjectRailItem__title']")

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
