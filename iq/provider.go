package iq

import (
	"iq-bot/core"
	
// 	"github.com/go-rod/rod"
// 	// "github.com/go-rod/rod/lib/input"
// 	// "time"
// 	// "fmt"
)

// type IqProvider struct {
// 	core.WebsiteLogin
// 	core.Connection
// }

type IqRequest struct {
	title string
	content string
	// author string
}

func GetRequests(connect core.Connection)(reqs []IqRequest){
	//scrape website
	elems:=GetTitles(connect)

	//get the content of each request
	for _, elem := range elems {
		content := GetContent(connect, elem)
		reqs = append(reqs, IqRequest{title: elem.MustText(), content: content})
	}
	return reqs
}
