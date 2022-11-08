package core

import (
	"errors"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	// "time"
	// "fmt"
)

// for holding information about a login
type WebsiteLogin struct {
	Url      string
	Username string
	Password string
}

// for holding information about a rod connection
type Connection struct {
	Browser *rod.Browser
	Page    *rod.Page
}

// should be able to work with most websites that use a login form
func SimpleLogin(connect Connection, login WebsiteLogin) {
	page := connect.Page
	//Race Condition: It will keep polling until one selector has found a match
	page.Race().Element("input[name='email']").MustHandle(func(e *rod.Element) {
		e.MustInput(login.Username).MustType(input.Enter)
	}).Element("input[name='username']").MustHandle(func(e *rod.Element) {
		e.MustInput(login.Username).MustType(input.Enter)
	}).MustDo()

	page.MustElement("input[name='password']").MustInput(login.Password).MustType(input.Enter)
}

func Connect(browser *rod.Browser, url string) (Connection) {
	page := browser.MustPage(url)
	return Connection{Browser: browser, Page: page}
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

	//login to the page
	page.Race().Element("input[name='email']").MustHandle(func(e *rod.Element) {
		e.MustInput(login.Username).MustType(input.Enter)
	}).Element("input[name='username']").MustHandle(func(e *rod.Element) {
		e.MustInput(login.Username).MustType(input.Enter)
	}).MustDo()

	page.MustElement("input[name='password']").MustInput(login.Password).MustType(input.Enter)

	//create connection object to return
	return Connection{Browser: browser, Page: page}, nil
}
