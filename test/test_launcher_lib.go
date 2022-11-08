package main

import (
    "github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
    "fmt"
)

func main() {
    launch := launcher.New().Bin("/Applications/Google Chrome.app/Contents/MacOS/Google Chrome").Headless(false).MustLaunch()
    browser := rod.New().ControlURL(launch).MustConnect()
    page := browser.MustPage("https://www.example.com")
    fmt.Println(page.MustElement("h1").MustText())
    browser.MustClose()
    // page := browser.MustPage("https://www.example.com")
    // fmt.Println(page.MustElement("h1").MustText())
    // fmt.Println("done")
}