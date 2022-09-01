package iq

import (
	"iq-bot/core"

	"github.com/go-rod/rod"
)

// type IqProxy interface {
// 	CreateProvider(browserString string) IqProvider
// 	Login(connect core.Connection, login core.WebsiteLogin)
// 	GetRequests(connect core.Connection) []IqRequest //defaults to 20 requests
// 	GetBatchRequests(connect core.Connection, number int) []IqRequest
// }

type IqProvider struct {
	browserString string
	core.Connection
	requests []IqRequest
}

type IqRequest struct {
	element *rod.Element
	title   string
	content string
	author  string
}

func CreateProvider(browserString string) (connect IqProvider) {
	//create a new connection
	connect = IqProvider{}
	connect.browserString = browserString
	connect.Connection = core.Connect(connect.browserString, "https://iq.aws.amazon.com/work/#/requests")
	return connect
}

func Login(connect core.Connection) {
	//load login information into memory
	cliEnv, err := core.LoadEnv()
	core.PrintIfErr(err)
	core.Success("environment : ", cliEnv)
	core.Login(connect, core.WebsiteLogin{cliEnv.Url, cliEnv.Username, cliEnv.Password})
	//(dont forget to manually enter 2fa)
}

func GetRequests(connect core.Connection) (reqs []IqRequest) {
	//scrape website
	elems := GetTitles(connect)

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
