package iq

import (
	"iq-bot/core"
	"github.com/go-rod/rod"
	// "github.com/go-rod/rod/lib/input"
	// "time"
	// "fmt"
)



func GetTitles(connect core.Connection) rod.Elements {
	elems := connect.Page.MustWaitLoad().MustElements("div[class^='ProjectRailItem__title']")
	return elems
}

func GetContent(connect core.Connection , elem *rod.Element) string{
		elem.MustClick()
		temp := connect.Page.MustWaitLoad().MustElement("div[class^='TextMessage__text']").MustText()
		return temp
}