package gateclient

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"runtime"

	"golang.org/x/oauth2"
)

const localWebServer = "localhost:8085"

func startWebServer() (codeCh chan string, err error) {
	listener, err := net.Listen("tcp", localWebServer)
	if err != nil {
		return nil, err
	}
	codeCh = make(chan string)

	go http.Serve(listener, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: add handle of `error`
		//error := r.FormValue("error")
		code := r.FormValue("code")
		codeCh <- code // send code to OAuth flow
		listener.Close()
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Received code: %v\r\nYou can now safely close this browser window.", code)
	}))

	return codeCh, nil
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(output func(string), config *oauth2.Config, authURL string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	codeCh, err := startWebServer()
	if err != nil {
		output("Unable to start a web server.")
		return nil, err
	}

	err = openURL(authURL)
	if err != nil {
		log.Fatalf("Unable to open authorization URL in web server: %v", err)
	} else {
		output("Your browser has been opened to an authorization URL.\n" +
			" This program will resume once authorization has been provided.\n")
		output(authURL)
	}

	// Wait for the web server to get the code.
	code := <-codeCh
	return exchangeToken(config, code, opts...)
}

// Exchange the authorization code for an access token
func exchangeToken(config *oauth2.Config, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	tok, err := config.Exchange(context.Background(), code, opts...)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve token %v", err)
	}
	return tok, nil
}

// openURL opens a browser window to the specified location.
// This code originally appeared at:
//   http://stackoverflow.com/questions/10377243/how-can-i-launch-a-process-that-is-not-a-file-in-go
func openURL(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("Cannot open URL %s on this platform", url)
	}
	return err
}
