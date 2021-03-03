package main

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/yeboahnanaosei/edith"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// showUsage prints to stdout details of how to use this tool
func showUsage() {
	usage := `
USAGE: edith ACTION RECIPIENT ITEM PAYLOAD

Send a one line text to Nana
 eg: edith send nana text "The text you want to send"

Send a file to nana
 eg: edith send nana file /path/to/file/you/want/to/send.txt

Send a multi-line text to Kwakye
 eg: edith send kwakye text
 Type or paste your multiline
 text here. The lines can be as
 many as you want

 Ctrl+D (Linux/Mac) or Ctrl+Z(Windows) to send

Send some credentials to Danny
 eg: edith send danny text
  Username: username
  Password: password

 Ctrl+D (Linux/Mac) or Ctrl+Z(Windows) to send

DETAILS:
ACTION
  ACTION is what you want edith to do
  value: send or get
   eg: edith send RECIPIENT ITEM PAYLOAD

RECIPIENT
  RECIPIENT is the person you are sending to
  value: name of a colleague eg. kwakye, nana etc.

   eg: edith ACTION nana ITEM PAYLOAD

ITEM
  ITEM is the thing you want to send either a text or file
  value: text or file

   eg: edith ACTION RECIPIENT text PAYLOAD
   eg: edith ACTION RECIPIENT file PAYLOAD

PAYLOAD
  PAYLOAD is the actual thing you are sending
  value:
    if ITEM is "file", then value for PAYLOAD is the path to the file

	eg. edith send nana file /path/to/the/file.ext

    if ITEM is "text", then value for PAYLOAD is the text you want to send

	eg. edith send nana text "The text I want to send"

    If you want to send a multiline text, don't supply payload. Just press
    enter and type your text. When you are done just hit the ENTER key
    followed by Ctrl+D on Linux/Mac or Ctrl+Z on Windows

    Notice no payload after 'text'
    eg. edith send nana text
    Type or paste your multi-line text here.
    You are free to type as many lines as you want

    Ctrl+D or Ctrl+Z // To finish your text
`
	fmt.Fprintf(os.Stdout, usage)
}

// performFirstRun takes the user's name and saves it in the config file
func performFirstRun(configFile io.Writer) (string, error) {
	fmt.Println("\nThis appears to be your first run. Let's get you setup")
	fmt.Print("Please enter your name: ")
	var name string
	fmt.Scanln(&name)

	if len(name) == 0 {
		return "", errors.New("no name supplied")
	}

	_, err := configFile.Write([]byte(fmt.Sprintf("name: %s", name)))
	if err != nil {
		return "", fmt.Errorf("could not write name to config file: %v", err)
	}
	fmt.Println()
	return name, nil
}

// getUserName gets the name of the user calling this tool.
// We need the name of the user to prepare a request to be sent to the server.
// If it is the first time the program is being called, performFirstRun() will be
// called
func getUserName() (string, error) {
	configFilePath := filepath.Join(os.Getenv("HOME"), ".edithrc")
	configFile, err := os.OpenFile(configFilePath, os.O_CREATE|os.O_RDWR, os.ModePerm)

	if os.IsNotExist(err) {
		return performFirstRun(configFile)
	}
	defer configFile.Close()

	finfo, err := configFile.Stat()
	if err != nil {
		return "", fmt.Errorf("could not stat config file: %v", err)
	}

	if finfo.Size() < 1 {
		return performFirstRun(configFile)
	}

	b, err := ioutil.ReadAll(configFile)
	if err != nil {
		return "", fmt.Errorf("could not read from config file: %v", err)
	}

	return string(bytes.Split(b, []byte(": "))[1]), nil
}

func getTextInput() ([]byte, error) {
	var input string
	fmt.Print("You can type or paste the text you want to send below:\n(Ctrl+D to send | Ctrl+C to cancel)\n\n")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		input += s.Text() + "\n"
	}

	if len(input) == 0 {
		return nil, errors.New("cannot continue with empty text. you need to supply some text")
	}

	return []byte(input), nil
}

// prepareRequest parses the args from the command line and prepares an appropriate
// request object to send to the server
func prepareRequest(args []string) (*edith.Request, error) {
	if len(args) < 3 {
		return nil, errors.New("not enough arguments supplied. at least three args must be provided")
	}

	// actions := []string{"send", "get"} // actions are the actions that edith can perform. edith can send or get items
	// items := []string{"text", "file"}  // items are the things that can be sent by edith. edith can send text or files

	sender, err := getUserName()
	if err != nil {
		return nil, fmt.Errorf("failed to get sender name: %v", err)
	}
	action := strings.ToLower(args[0])
	recipient := strings.ToLower(args[1])
	item := strings.ToLower(args[2])

	if action != "send" && action != "get" {
		return nil, fmt.Errorf("unknown action '%s', expected one of [send, get]", action)
	}

	if recipient == "" {
		return nil, errors.New("you did not supply the name of your recipient. recipient cannot be empty")
	}

	if item != "file" && item != "text" {
		return nil, fmt.Errorf("unknown item '%s'. expected one of [file, text]", item)
	}

	req := &edith.Request{}
	switch action {
	case "send":
		// Handle if item type it text
		if item == "text" && len(args) > 3 {
			req.Body = []byte(strings.Join(args[3:], " "))
			if len(req.Body) == 0 {
				return nil, errors.New("cannot continue with empty text. you need to supply some text")
			}
			req.Type = "text"
		} else if item == "text" && len(args) == 3 {
			req.Body, err = getTextInput()
			if err != nil {
				return nil, err
			}
			if len(req.Body) == 0 {
				return nil, errors.New("cannot continue with empty text. you need to supply some text")
			}
			req.Type = "text"
		}

		// Handle if user wants to send a file
		if item == "file" && len(args) == 3 {
			return nil, errors.New("you are attempting to send a file but you did not supply the path to the file")
		} else if item == "file" && len(args) > 3 {
			abs, err := filepath.Abs(args[3])
			if err != nil {
				return nil, fmt.Errorf("could not determine the absolute path to the file you supplied %s: %v", args[3], err)
			}
			req.Body, err = ioutil.ReadFile(abs)
			if err != nil {
				return nil, fmt.Errorf("could not get the contents of the file you supplied %s: %v", args[3], err)
			}
			if len(req.Body) == 0 {
				return nil, errors.New("file appears to be empty")
			}
			req.Filename = args[3]
			req.Type = "file"
		}

		req.Sender, req.Recipient = sender, recipient
	case "get":
		req.Sender, req.Recipient, req.Type = recipient, sender, item
	}
	return req, nil
}

// makeRequest makes the request to the server
func makeRequest(ctx context.Context, action string, client edith.EdithClient, request *edith.Request) (interface{}, error) {
	action = strings.ToLower(action)
	if action == "send" {
		switch request.Type {
		case "file":
			return client.SendFile(ctx, request)

		case "text":
			return client.SendText(ctx, request)
		default:
			return nil, errors.New("unknown item type: " + request.Type)
		}
	}

	if action == "get" {
		switch request.Type {
		case "file":
			return client.GetFile(ctx, request)

		case "text":
			return client.GetText(ctx, request)
		default:
			return nil, errors.New("unknown item type: " + request.Type)
		}
	}

	return nil, errors.New("unknown action type: " + action)

}

// This can be changed using ldflags when building for release
var serverAddr = "localhost:54920"

// createClient creates the edith grpc client for communication with the server
func createClient() (edith.EdithClient, error) {
	con, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return edith.NewEdithClient(con), nil
}

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "edith: not enough arguments supplied\n")
		showUsage()
		os.Exit(1)
	}

	request, err := prepareRequest(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "edith: error preparing request: %v\n", err)
		os.Exit(1)
	}

	client, err := createClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "edith: error creating edith client: %v\n", err)
		os.Exit(1)
	}

	res, err := makeRequest(context.Background(), os.Args[1], client, request)
	st, ok := status.FromError(err)
	if ok && st.Err() != nil {
		fmt.Fprintf(os.Stderr, "edith: %v\n", st.Message())
		os.Exit(1)
	}

	switch r := res.(type) {
	case *edith.RequestItems:
		itemsLen := len(r.Texts)
		if itemsLen == 0 {
			fmt.Fprintf(os.Stdout, "\nYou have no items from %s\n", strings.Title(os.Args[2]))
			os.Exit(0)
		}

		fmt.Fprintf(os.Stdout, "\nLast 5 texts from %s:\n\n", strings.Title(os.Args[2]))

		for x := itemsLen; x > 0; x-- {
			fmt.Fprintf(os.Stdout, "%d.\n%s\n\n", x, r.Texts[x-1].Body)
		}
		os.Exit(0)

	case *edith.Response:
		fmt.Fprintf(os.Stdout, "edith: %v\n", r.Msg)
		os.Exit(0)

	default:
		fmt.Fprintf(os.Stderr, "edith: unknown type of res: %T", res)
		os.Exit(1)
	}
}
