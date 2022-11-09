package main

import (
	// "database/sql"
	// "fmt"
	"fmt"
	"iq-bot/cli"
	"iq-bot/core"
	"iq-bot/iq"
	"os"

	// "reflect"
	// "strconv"
	"strings"
	"time"

	// "github.com/go-rod/rod"
	"github.com/manifoldco/promptui"
	// "errors"
	// "reflect"
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

func Bootstrap(iq.IqProvider) (p iq.IqProvider, err error) {
	cli.Welcome()

	//load login information into memory
	p.AwsEnv, err = core.LoadEnv()
	cli.Success("environment : ", p.AwsEnv)
	cli.PrintIfErr(err)

	cli.Success("len(os.Args) : ", len(os.Args))
	if len(os.Args) > 2 {
		browser := core.Manual(os.Args[2])
		p.Connection.Browser = browser
		cli.Success("p.Connection : ", p.Connection)
		p.LoggedIn = true
		p.Connection.Page = p.Connection.Browser.MustPage("https://iq.aws.amazon.com/work/#/requests")
		// p = NavIq(p)

		return p, err
	}

	if !p.LoggedIn {
		p, err = Authenticate(p)
		cli.PrintIfErr(err)
		p = NavIq(p)


	}
	cli.Success("p.LoggedIn : ", p.LoggedIn)

	return p, err
}

// func CreateBrowser() (browser *rod.Browser) {
// 	browser = rod.New().MustConnect()
// 	// if path, exists := launcher.LookPath(); exists {
// 	// 	u := launcher.New().Bin(path).MustLaunch()
// 	// 	rod.New().ControlURL(u).MustConnect()
// 	// }
// 	return browser
// }

func Authenticate(p iq.IqProvider) (iq.IqProvider, error) {

	//declare error
	var err error

	cli.Success("p.Connection before: ", p.Connection)
	p.Connection = core.Connect(p.Connection.Browser, p.AwsEnv.Url)
	cli.Success("p.Connection after: ", p.Connection)

	//login to aws
	p.Connection, err = core.SimpleLogin(p.Connection, core.WebsiteLogin{p.AwsEnv.Url, p.AwsEnv.Username, p.AwsEnv.Password})
	//wait for 2fa - this is a hack for now, need to remove
	cli.Success("p.Connection: ", p.Connection)
	cli.PrintIfErr(err)
	cli.Success("...waiting on 2FA... (return to browser)")
	//authentication boolean
	// p.LoggedIn = true
	p.LoggedIn = true
	cli.Success("p.LoggedIn : ", p.LoggedIn)
	return p, err

}

func NavIq(iq.IqProvider) (p iq.IqProvider) {
	// p.Connection.Page.MustNavigate("https://iq.aws.amazon.com/work/#/requests")
	p.Connection.Page = p.Connection.Browser.MustPage("https://iq.aws.amazon.com/work/#/requests")
	return p
}

func Execute(p iq.IqProvider) {
	cli.Success("iq.IqProvider : ", p)

	options := []cli.PromptOptions{
		{
			Label: "Exit CLI",
			Key:   0,
		},
		{
			Label: "Reauthenticate",
			Key:   1,
		},
		{
			Label: "Navigate to IQ",
			Key:   2,
		},

		{
			Label: "Get Requests",
			Key:   3,
		},

		// {
		// 	Label: "Send Message",
		// 	Key:   3,
		// },
		// {
		// 	Label: "Send Bulk Message",
		// 	Key:   4,
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
	case 1:
		cli.Success("Authenticate")
		p, _ = Authenticate(p)
		// iq.GetRequests(p.Connection)
	case 2:
		cli.Success("Navigate to IQ")
		p = NavIq(p)
	case 3:
		cli.Success("Get Requests")
		p = GetRequests(p)
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

// func GetBrowser() (browser *rod.Browser) {
// 	browserString := cli.PromptGetInput(cli.PromptContent{
// 				Label: "Please enter the browser string",
// 			})
// 	browser = core.Manual(browserString)
// 	return browser
// }

// func Dash(u *rod.Browser){

// 		// fmt.Println("Arg length is %d", argLength)
// 		// u := os.Args[2]

// 		//navigate to IQ
// 		browser = core.Manual(u)
// 		connect := core.Connect(browser, "https://iq.aws.amazon.com/work/#/requests")
// 		cli.Success("connection : ", connect)

// 		//takes a second, I guess
// 		time.Sleep(time.Second * 2)

// 		reqs := iq.GetRequests(connect)

// 		for _, req := range reqs {
// 			cli.Success("request : ", req)
// 			//get value of title and content
// 			title := reflect.ValueOf(req).FieldByName("title").String()
// 			content := reflect.ValueOf(req).FieldByName("content").String()
// 			author := reflect.ValueOf(req).FieldByName("author").String()
// 			cli.Success("title : ", title)
// 			cli.Success("author : ", author)
// 			cli.Success("content : ", content)
// 		}

// 		cli.Success("# of requests: ", strconv.Itoa(len(reqs)))

// }

// go core.Manager()
// time.Sleep(1 * time.Second)
// go core.Remote()
// func Login(){
// 	cliEnv, err := core.LoadEnv()
// 	cli.PrintIfErr(err)
// 	cli.Success("environment : ", cliEnv)

// 	//connect to aws
// 	connect := core.Connect(u, cliEnv.Url)
// 	cli.PrintIfErr(err)
// 	cli.Success("connection : ", connect)

// 	//enter login credentials
// 	core.Login(connect, core.WebsiteLogin{cliEnv.Url, cliEnv.Username, cliEnv.Password})
// }

// func promptGetInput(pc promptContent) string {
// 	validate := func(input string) error {
// 		if len(input) <= 0 {
// 			return errors.New(pc.errorMsg)
// 		}
// 		return nil
// 	}
// 	prompt := promptui.Prompt{
// 		Label: pc.Label,
// 		// Templates: templates,
// 		Validate: validate,
// 	}
// 	result, err := prompt.Run()
// 	if err != nil {
// 		fmt.Printf("Prompt failed %v\n", err)
// 		os.Exit(1)
// 	}
// 	fmt.Printf("Input: %s\n", result)
// 	return result
// }

// func GetBatchRequestNumber(connect core.Connection) (reqs []iq.IqRequest) {
// 	numberRequests := promptGetInput(promptContent{
// 		Label: "Enter the number of requests you want to scrape",
// 	})

// 	num, _ := strconv.Atoi(numberRequests)

// 	reqs = iq.GetBatchRequests(connect, num)
// 	cli.Success("reqs : ", reqs)
// 	return reqs
// }

// func PrintRequests(reqs []iq.IqRequest) {
// 	for _, req := range reqs {
// 		cli.Success("Request:", reflect.ValueOf(req).FieldByName("title").String())
// 	}
// }

// func SendMessage(connect core.Connection, templates *promptui.SelectTemplates, searcher func(input string, index int) bool, reqs []iq.IqRequest) {

// 	// options := []iq.IqRequest{reqs...}
// 	options := []promptOptions{}
// 	for _, req := range reqs {
// 		cli.Success("Request:", reflect.ValueOf(req).FieldByName("title").String())
// 		// options = append(options, promptOptions{
// 		// 	Label: req.title,
// 		// 	Key:   int64(req.id),
// 		// })
// 	}

// 	prompt := promptui.Select{
// 		Label:     "Welcome to IQ-Bot. Please choose an option: ",
// 		Items:     options,
// 		Templates: templates,
// 		// Size:      4,
// 		Searcher: searcher,
// 	}

// 	i, _, err := prompt.Run()

// 	if err != nil {
// 		fmt.Printf("Prompt failed %v\n", err)
// 		main()
// 		// return ACloudEnv{}, err
// 	}

// 	fmt.Printf("You choose request %d: %s\n", i+1, options[i].Label)

// 	// iq.InsertMessage(connect, req.element, msg)

// }

/*
*

  - Execute the program
    func GetSandboxCreds(cliEnv core.ACloudEnv, p *acloud.ACloudProvider) (acloud.ACloudProvider, error) {

    //connect to website
    connect, err := core.Login(core.WebsiteLogin{Url: cliEnv.Url, Username: cliEnv.Username, Password: cliEnv.Password})
    cli.PrintIfErr(err)
    cli.Success("Connection Successful: ", connect)
    p.Connection = connect

    //scrape credentials
    elems, err := acloud.Sandbox(p.Connection, cliEnv.Download_key)
    cli.PrintIfErr(err)
    // cli.Success("rod html elements : ", elems)

    //copy credentials to clipboard
    creds, err := acloud.CopyHtml(elems)
    cli.PrintIfErr(err)
    // cli.Success("credentials : ", creds)
    p.SandboxCredential = creds

    //DISPLAY WITH COLORS PROMINENTLY TO THE USER
    acloud.DisplayCreds(creds)

    return *p, err
    }
*/

// go core.Manager()
// go core.Remote()

// _, err := exec.Command("/bin/sh", "./scripts/frame.sh").Output()
// if err != nil {
// 	panic(err)
// }
// cli.Success("Script executed successfully")

// var p iq.IQProvider
// cli.Success("getting iq provider login...")

// if len(os.Args) > 1 {
// 	cli.Success("setting args to env...")
// 	env, err :=core.ArgEnv()
// 	cli.PrintIfErr(err)
// 	p.IQEnv = env
// 	p = bootstrap(p)
// 	Execute(p)
// } else {
// 	env, err := cli.GetEnv(".env")
// 	cli.PrintIfErr(err)
// 	if err != nil {
// 		cli.Error("Error: .env file not found")
// 		env = core.Env()
// 	}
// 	cli.Success("env : ", env)
// 	p.IQEnv = env
// 	p = bootstrap(p)
// 	Execute(p)
// }
