package iq

import (
	"fmt"
	"github.com/go-rod/rod"
	"iq-bot/cli"
	"iq-bot/core"
	// "strconv"
)

// type IqProxy interface {
// 	CreateProvider(browserString string) IqProvider
// 	Login(connect core.Connection, login core.WebsiteLogin)
// 	GetRequests(connect core.Connection) []IqRequest //defaults to 20 requests
// 	GetBatchRequests(connect core.Connection, number int) []IqRequest
// }

type IqProvider struct {
	LoggedIn bool
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
	date    string
}

func GetRequests(connect core.Connection) (reqs []IqRequest) {
	//scrape website
	elems := GetTitles(connect)
	fmt.Println("elems : ", elems)

	//get the content of each request
	for _, elem := range elems {
		content := GetContent(connect, elem)
		author := GetAuthor(connect, elem)

		reqs = append(reqs, IqRequest{
			element: elem,
			title:   elem.MustText(),
			content: content,
			author:  author,
		})
	}
	return reqs
}

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
	// elems := connect.Page.MustWaitLoad().MustElements("div[class^='content ProjectRailItem']")
	//log the number of elements
	// fmt.Println("Number of elements : ", len(elems))

	//log number of titles
	fmt.Println("Number of titles : ", len(connect.Page.MustWaitLoad().MustElements("div[class^='ProjectRailItem__title']")))
	//log number of dates
	fmt.Println("Number of dates : ", len(connect.Page.MustWaitLoad().MustElements("small[class^='ProjectRailItem__date']")))
	//log number of labels
	// fmt.Println("Number of labels : ", len(connect.Page.MustWaitLoad().MustElements("div[class^='ProjectRailItem__label']")))
	//log number of budgets
	// fmt.Println("Number of budgets : ", len(connect.Page.MustWaitLoad().MustElements("div[class^='ProjectRailItem__budget']")))

	dates := connect.Page.MustWaitLoad().MustElements("small[class^='ProjectRailItem__date']")
	titles := connect.Page.MustWaitLoad().MustElements("div[class^='ProjectRailItem__title']")
	// labels := connect.Page.MustWaitLoad().MustElements("div[class^='ProjectRailItem__label']")
	// budgets := connect.Page.MustWaitLoad().MustElements("div[class^='ProjectRailItem__budget']")

	//loop over length of elems
	for i := 0; i < len(titles); i++ {
		// var tempElem *rod.Element

		//declare temporary vars to hold values
		var tempTitle string
		var tempContent string
		var tempAuthor string
		// var tempBudget string
		// var tempLabel string
		var tempDate string

		//check each element for nil

		// if error := elems[i]; error != nil {
		// 	tempElem = elems[i]
		// }
		if error := titles[i]; error != nil {
			tempTitle = titles[i].MustText()
		}
		// if error := labels[i]; error != nil {
		// 	tempLabel = labels[i].MustText()
		// }
		// if error := budgets[i]; error != nil {
		// 	tempBudget = budgets[i].MustText()
		// }
		if error := dates[i]; error != nil {
			tempDate = dates[i].MustText()
		}

		//get content and author
		// tempContent, tempAuthor = GetRequestDetails(connect, tempElem)
		// use error handling to check for nil
		if error, _ := GetRequestDetails(connect, titles[i]); error != "" {
			tempContent, tempAuthor = GetRequestDetails(connect, titles[i])
		} else {
			continue
		}

		reqs = append(reqs, IqRequest{
			title:   tempTitle,
			date:    tempDate,
			// label:   tempLabel,
			// budget:  tempBudget,
			content: tempContent,
			author:  tempAuthor,
		})
	}

	//for each elem, get the text of the header, title, label, and budget
	//create a new IqRequest object with the text of each element
	//append the new IqRequest object to the array of IqRequest objects
	//return the array of IqRequest objects

	return reqs
}

func GetRequestDetails(connect core.Connection, elem *rod.Element) (content string, author string) {
	//get the title, content, author, budget, label, and header
	elem.MustClick()
	content = connect.Page.MustWaitLoad().MustElement("div[class^='TextMessage__text']").MustText()
	author = connect.Page.MustWaitLoad().MustElement("span[class^='author']").MustText()

	return content, author
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
	for _, req := range reqs {
		InsertMessage(connect, req.element, "Hello, I am a bot. I am responding to your request.")
	}
}
