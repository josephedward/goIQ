package core

import (
	// "errors"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	// "time"
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
func Login(connect Connection, login WebsiteLogin) {
	page := connect.Page
	//Race Condition: It will keep polling until one selector has found a match
	page.Race().Element("input[name='email']").MustHandle(func(e *rod.Element) {
		e.MustInput(login.Username).MustType(input.Enter)
	}).Element("input[name='username']").MustHandle(func(e *rod.Element) {
		e.MustInput(login.Username).MustType(input.Enter)
	}).MustDo()

	page.MustElement("input[name='password']").MustInput(login.Password).MustType(input.Enter)
}

func Connect(u, url string) (Connection, error) {
	browser := Manual(u)
	page := rod.New().ControlURL(u).MustConnect().MustPage(url)
	return Connection{Browser: browser, Page: page}, nil
}
