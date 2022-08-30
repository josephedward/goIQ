package core

import (
	"os"
	"os/exec"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/utils"
	"github.com/ysmood/leakless"
)

func CustomLaunch() *rod.Browser{
	// get the browser executable path
	path := launcher.NewBrowser().MustGet()

	// use the FormatArgs to construct args, this line is optional, you can construct the args manually
	args := launcher.New().FormatArgs()

	var cmd *exec.Cmd
	if true { // decide whether to use leakless or not
		cmd = leakless.New().Command(path, args...)
	} else {
		cmd = exec.Command(path, args...)
	}

	parser := launcher.NewURLParser()
	cmd.Stderr = parser
	utils.E(cmd.Start())
	u := launcher.MustResolveURL(<-parser.URL)

	return rod.New().ControlURL(u).MustConnect()
}


func UseSystemBrowser() *rod.Browser {
	if path, exists := launcher.LookPath(); exists {
		u := launcher.New().Bin(path).MustLaunch()
		return rod.New().ControlURL(u).MustConnect()
	}
	return nil
}


func BrowserCliOutput() *rod.Browser{
	// Pipe the browser stderr and stdout to os.Stdout .
	u := launcher.New().Logger(os.Stdout).MustLaunch()
	return rod.New().ControlURL(u).MustConnect()
}

func Manual(u string) *rod.Browser{
	return rod.New().ControlURL(u).MustConnect()
}