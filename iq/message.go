package iq

import (
	"iq-bot/core"
	"github.com/go-rod/rod"
)


func InsertMessage(connect core.Connection , elem *rod.Element, message string) {
	elem.MustClick()
	textarea := connect.Page.MustWaitLoad().MustElement("textarea[id^='initialResponse']")
	textarea.MustSelectAllText().MustInput("")
	textarea.MustInput(message)
	core.Success("Message inserted")
}


func DraftIntro()(){}

func AnswerAllRequests(connect core.Connection , reqs []IqRequest) {
	// for _, req := range reqs {
		
	// 	// AnswerRequest(connect, req)
	// }
}