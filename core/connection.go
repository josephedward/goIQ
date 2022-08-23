package core

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	//launcher
	// "github.com/go-rod/rod/lib/launcher"
	"errors"
)

type WebsiteLogin struct {
	Url      string
	Username string
	Password string
}

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

	//login to the page
	page.Race().Element("input[name='email']").MustHandle(func(e *rod.Element) {
		e.MustInput(login.Username).MustType(input.Enter)
	}).Element("input[name='username']").MustHandle(func(e *rod.Element) {
		e.MustInput(login.Username).MustType(input.Enter)
	}).MustDo()
	

	page.MustElement("input[name='password']").MustInput(login.Password).MustType(input.Enter)

	// page.MustWaitLoad().MustElementR("input", "Search").MustInput("iq").MustType(input.Enter)


	//create connection object to return
	return Connection{Browser: browser, Page: page}, nil
}


func Visit(url string){
//  (Connection, error)
  
	// Launch a new browser with default options, and connect to it.
	// browser := rod.New().MustConnect()

	// // Create a new page
	// page := browser.MustPage(url)

	u:="ws://127.0.0.1:9222/devtools/browser/6a023f38-873e-4aa1-b5f3-45353ff3a72b"
	rod.New().ControlURL(u).MustConnect().MustPage(url)
	
	// u := launcher.New().Bin("/Applications/Google Chrome.app/Contents/MacOS/Google Chrome").MustLaunch()
    // rod.New().ControlURL(u).MustConnect().MustPage(url)


	// //if browser is nil, page is nil
	// if browser == nil || page == nil {
	// 	return Connection{}, errors.New("browser or page is nil")
	// }

	// //create connection object to return
	// return Connection{Browser: browser, Page: page}, nil
}