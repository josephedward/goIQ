package main

import (
	// "database/sql"
	// "fmt"
	"fmt"
	"iq-bot/cli"
	"iq-bot/core"
	"iq-bot/iq"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"github.com/go-rod/rod"
	"github.com/manifoldco/promptui"
	// "errors"
)

func main() {

	var p iq.IqProvider
	p, err := Bootstrap(p)
	cli.PrintIfErr(err)
	exit := false
	for !exit {
		Execute(p)
	}

}

func Bootstrap(iq.IqProvider) (p iq.IqProvider, err error) {
	cli.Welcome()

	//cli arguments
	browserString := os.Args[2]
	if len(browserString) != 0 {
		cli.Success("Browser String : ", browserString)
		b := core.Manual(browserString)
		cli.Success("Browser : ", b)
		Dash(b)
	}

	

	
	//load login information into memory
	p.AwsEnv, err = core.LoadEnv()
	cli.Success("environment : ", p.AwsEnv)
	cli.PrintIfErr(err)

	//login to aws
	p.Connection, err = core.Login(core.WebsiteLogin{p.AwsEnv.Url, p.AwsEnv.Username, p.AwsEnv.Password})
	cli.Success("logged in : ", p.Connection)
	cli.PrintIfErr(err)
	cli.Success("...waiting on 2FA... (return to browser)")

	// iqPage := p.Connection.Browser.MustPage("https://iq.aws.amazon.com/work/#/requests")
	// p.Connection.Page = iqPage;
	// cli.Success("iqPage : ", iqPage)

	return p, err
}

func Execute(p iq.IqProvider) {
	cli.Success("iq.IqProvider : ", p)

	options := []cli.PromptOptions{
		{
			Label: "Exit CLI",
			Key:   0,
		},
		{
			Label: "Get Requests",
			Key:   1,
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
		cli.Success("Get Requests")
		GetRequests(p)
		// iq.GetRequests(p.Connection)
	case 2:
		// 	cli.Success("Send Message")
		// case 3:
		// 	cli.Success("Send Bulk Message")
	}

	Execute(p)
}

func GetRequests(p iq.IqProvider) {
	p.Connection.Page = p.Connection.Page.MustNavigate("https://iq.aws.amazon.com/work/#/requests")
	//stall for page to load
	// p.Connection.Page.MustWaitLoad()
	//takes a second, I guess
	time.Sleep(time.Second * 2)

	iq.GetRequests(p.Connection)

}

func GetBrowser() (browser *rod.Browser) {
	browserString := cli.PromptGetInput(cli.PromptContent{
				Label: "Please enter the browser string",
			})
	browser = core.Manual(browserString) 
	return browser
}

func Dash(u *rod.Browser){

		// fmt.Println("Arg length is %d", argLength)
		// u := os.Args[2]
	
		//navigate to IQ
		browser = core.Manual(u) 
		connect := core.Connect(browser, "https://iq.aws.amazon.com/work/#/requests")
		cli.Success("connection : ", connect)
	
		//takes a second, I guess
		time.Sleep(time.Second * 2)
	
		reqs := iq.GetRequests(connect)
	
		for _, req := range reqs {
			cli.Success("request : ", req)
			//get value of title and content
			title := reflect.ValueOf(req).FieldByName("title").String()
			content := reflect.ValueOf(req).FieldByName("content").String()
			author := reflect.ValueOf(req).FieldByName("author").String()
			cli.Success("title : ", title)
			cli.Success("author : ", author)
			cli.Success("content : ", content)
		}
	
		cli.Success("# of requests: ", strconv.Itoa(len(reqs)))
	
	
}

// go core.Manager()
// time.Sleep(1 * time.Second)
// go core.Remote()
// func Login(){
// 	cliEnv, err := core.LoadEnv()
// 	core.PrintIfErr(err)
// 	cli.Success("environment : ", cliEnv)

// 	//connect to aws
// 	connect := core.Connect(u, cliEnv.Url)
// 	core.PrintIfErr(err)
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
