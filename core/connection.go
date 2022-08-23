package core

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	// "github.com/go-rod/rod/lib/launcher"
	"errors"
	"iq-bot/core/browser"
	"time"
)

//for holding information about a login
type WebsiteLogin struct {
	Url      string
	Username string
	Password string
}

//for holding information about a rod connection
type Connection struct {
	Browser *rod.Browser
	Page    *rod.Page
}

func Login(login WebsiteLogin) (Connection, error) {
	// Launch a new browser with default options, and connect to it.
	browser := rod.New().MustConnect()

	// Create a new page
	page := browser.MustPage(login.Url)
    
	//if browser is nil, page is nil
	if browser == nil || page == nil {
		return Connection{}, errors.New("browser or page is nil")
	}

	//Race Condition: It will keep polling until one selector has found a match
	page.Race().Element("input[name='email']").MustHandle(func(e *rod.Element) {
		e.MustInput(login.Username).MustType(input.Enter)
	}).Element("input[name='username']").MustHandle(func(e *rod.Element) {
		e.MustInput(login.Username).MustType(input.Enter)
	}).MustDo()
	

	page.MustElement("input[name='password']").MustInput(login.Password).MustType(input.Enter)

	//create connection object to return
	return Connection{Browser: browser, Page: page}, nil
}

func VisitManual(u, url string) (Connection, error){
	browser := browser.Manual(u)
	page :=rod.New().ControlURL(u).MustConnect().MustPage(url)
	time.Sleep(5 * time.Second)
	return Connection{Browser: browser, Page: page}, nil	
}
