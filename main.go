package main

import (
	"encoding/csv"
	"fmt"
	"github.com/manifoldco/promptui"
	"iq-bot/cli"
	"iq-bot/core"
	"iq-bot/iq"
	"os"
	"strings"
	"time"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	var p iq.IqProvider
	cli.Success("p.LoggedIn : ", p.LoggedIn)
	p, err := Bootstrap(p)
	cli.Success("IqProvider after Bootstrap: ", p)
	cli.PrintIfErr(err)
	exit := false
	for !exit {
		Execute(p)
	}
}

//function for testing current url to see if we are logged in
func CheckCurrentUrl(p iq.IqProvider) (string){
	url := p.Connection.Page.MustInfo().URL
	fmt.Println("url : ", url)
	return url 
}


func Bootstrap(iq.IqProvider) (p iq.IqProvider, err error) {
	cli.Welcome()

	// //load login information into memory
	// p.AwsEnv, err = core.LoadEnv()
	// cli.Success("environment : ", p.AwsEnv)
	// cli.PrintIfErr(err)

	cli.Success("len(os.Args) : ", len(os.Args))
	if len(os.Args) > 2 {
		// browser := core.Manual(os.Args[2])
		// p.Connection.Browser = browser
		p, err = ConnectBrowser(p)
	} else {
		// p = NavIq(p)
		p.Connection.Browser = core.BrowserCliOutput()
		// p, err = ConnectBrowser(p)
	}

	// if p.Connection.Page == nil {
	// 	p, err = Authenticate(p)
	// }

	p.Connection.Page = p.Connection.Browser.MustPage("https://iq.aws.amazon.com/work/#/requests")
	cli.Success("p : ", p)
	cli.PrintIfErr(err)
	return p, err
}



func ConnectBrowser(p iq.IqProvider) (iq.IqProvider, error) {
	u := launcher.MustResolveURL("")
	browser := rod.New().ControlURL(u).MustConnect()
	p.Connection.Browser = browser
	return p, nil
}


func NavIq(iq.IqProvider) (p iq.IqProvider) {
	// p.Connection.Page.MustNavigate("https://iq.aws.amazon.com/work/#/requests")
	//log current url 
	CheckCurrentUrl(p)
	// var err error
	// p, err = ConnectBrowser(p)
	// cli.PrintIfErr(err)

	p.Connection.Page = p.Connection.Browser.MustPage("https://iq.aws.amazon.com/work/#/requests")
	p.Connection.Page.MustNavigate("https://iq.aws.amazon.com/work/#/requests")
	return p
}

func Execute(p iq.IqProvider) {
	// cli.Success("iq.IqProvider : ", p)

	options := []cli.PromptOptions{
		{
			Label: "Exit CLI",
			Key:   0,
		},
		// {
		// 	Label: "Authenticate",
		// 	Key:   1,
		// },
		{
			Label: "Navigate to IQ",
			Key:   2,
		},
		{
			Label: "Get Requests",
			Key:   3,
		},
		{
			Label: "Display Requests",
			Key:   4,
		},
		{
			Label: "Display IQ Provider",
			Key:   5,
		},
		{
			Label: "Get # of Requests",
			Key:   6,
		},
		{
			Label: "Write CSV",
			Key:   7,
		},

		// {
		// 	Label: "Send Message",
		// },

	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Label | yellow }} ",
		Inactive: "  {{ .Label | cyan }} ",
		Selected: "\U0001FAD1 {{ .Label | green | cyan }}",
	}

	searcher := func(input string, index int) bool {
		option := options[index]
		name := strings.Replace(strings.ToLower(option.Label), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Welcome to IQ-Bot. Please choose an option: ",
		Items:     options,
		Templates: templates,
		// Size:      4,
		Searcher: searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
		// return ACloudEnv{}, err
	}

	fmt.Printf("You choose number %d: %s\n", i+1, options[i].Label)

	switch options[i].Key {
	case 0:
		os.Exit(0)
	// case 1:
	// 	cli.Success("Authenticate")
		// p, _ = Authenticate(p)
		// iq.GetRequests(p.Connection)
	case 2:
		cli.Success("Navigate to IQ")
		p = NavIq(p)
	case 3:
		cli.Success("Get Requests")
		p = GetRequests(p)
	case 4:
		cli.Success("Display Requests")
		iq.DisplayRequests(p.Requests)
	case 5:
		cli.Success("Display IQ Provider")
		cli.Success("p : ", p)
	case 6:
		cli.Success("Get # of Requests")
		p.Requests = iq.GetBatchRequests(p.Connection, 5)
	case 7:
		cli.Success("Write CSV")
		WriteCsv(p)
	}

	Execute(p)
}

func GetRequests(p iq.IqProvider) iq.IqProvider {

	//takes a second, I guess
	time.Sleep(time.Second * 2)
	cli.Success("p : ", p)

	//get all requests
	requests := iq.GetElements(p.Connection)
	cli.Success("requests : ", requests)
	p.Requests = requests
	return p
}

func WriteCsv(p iq.IqProvider) {

	filename := fmt.Sprintf("./request-data_%s.csv", time.Now())

	f, e := os.Create(filename)
	if e != nil {
		fmt.Println(e)
	}

	writer := csv.NewWriter(f)
	var data = [][]string{
		{"title ", "content", "author", "budget", "label", "date"},
	}

	for _, v := range p.Requests {
		data = append(data, []string{v.Title, v.Content, v.Author, v.Budget, v.Label, v.Date})
	}

	e = writer.WriteAll(data)
	if e != nil {
		fmt.Println(e)
	}
}
