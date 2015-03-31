package stelligent

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

const HOST_URL = "http://52.4.210.206/"
const SEARCH_STRING = "<h2>Automation for the People</h2>"

func Test_HttpService(t *testing.T) {
	response, err := http.Get(HOST_URL)
	if err != nil {
		t.Error("HTTP GET Error: %v\n", err)
	}

	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error("HTTP GET Read Error: %v", err)
	}

	if !strings.Contains(string(content), SEARCH_STRING) {
		t.Error("HTTP content doesn't match file")
	}
}
