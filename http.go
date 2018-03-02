package proxmox

import (
	"io"
	"net/http"
	"fmt"
	netURL "net/url"
	"strings"
	"encoding/json"
	"io/ioutil"
)

func Post(path string, body io.Reader) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", server, path)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return &http.Response{}, err
	}
	return client.Do(req)
}

func Get(path string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", server, path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &http.Response{}, err
	}
	return client.Do(req)
}

func generateRequestBody(data map[string]string) *strings.Reader {
	body := netURL.Values{}

	for key, value := range data {
		body.Set(key, value)
	}

	return strings.NewReader(body.Encode())
}

func unmarshalResponse(r *http.Response, target interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, target)
}

func setCookie(ticket string) {
	cookie = &http.Cookie{
		Name: "PVEAuthCookie",
		Value: ticket,
	}
}