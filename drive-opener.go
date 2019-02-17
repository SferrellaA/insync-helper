package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

/*************
TODO List
- Handle multiple accounts
- Open with different browsers
- remove insync branding?
- config file that is read from
- handle browser failing to open
- handle opening multiple files?
- associate file formats with this
- gracefully handle non insync, non gogle files
*************/

// errFail gives a notify-send message when there is an error
func errFail(err error) {
	if nil == err {
		return
	}

	cmd := exec.Command("notify-send", "Insync Error", err.Error())
	cmd.Run()
	log.Fatal(err)
}

// insync mirrors the .json format that InsyncHQ uses
type insync struct {
	Url     string `json:"url"`
	Account string `json:"account_email"`
	File    string `json:"file_id"`
}

func main() {
	// Read the raw data from the file
	bytes, err := ioutil.ReadFile(os.Args[1])
	errFail(err)

	// Unmarshall the raw json data
	var data insync
	err = json.Unmarshal(bytes, &data)
	errFail(err)

	// Open Chrome with appropriate profile and file
	cmd := exec.Command("chromium", "--profile-directory Default", data.Url)
	err = cmd.Start()
	errFail(err)
}
