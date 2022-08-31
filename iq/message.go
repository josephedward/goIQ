package iq

import (
	"iq-bot/core"
	// "iq-bot/iq"
	// "os"

	"github.com/go-rod/rod"
	// "github.com/go-rod/rod"
	// // "github.com/go-rod/rod/lib/input"
	// // "time"
	// // "fmt"
)


func InsertMessage(connect core.Connection , elem *rod.Element, message string) {
	elem.MustClick()
	connect.Page.MustWaitLoad().MustElement("textarea[id^='initialResponse']").MustInput(message)
	core.Success("Message inserted")
}