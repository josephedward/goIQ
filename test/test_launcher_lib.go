package main

import (
    "github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
    u := launcher.New().Bin("/Applications/Google Chrome.app/Contents/MacOS/Google Chrome").MustLaunch()
    rod.New().ControlURL(u).MustConnect().MustPage("https://example.com")
}