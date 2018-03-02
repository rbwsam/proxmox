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

func get(path string) (*http.Response, error) {
	return send("GET", path, nil)
}

func post(path string, body io.Reader) (*http.Response, error) {
	return send("POST", path, body)
}

func put(path string, body io.Reader) (*http.Response, error) {
	return send("PUT", path, body)
}

func del(path string) (*http.Response, error) {
	return send("DELETE", path, nil)
}

func send(method, path string, body io.Reader) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", server, path)
	req, err := http.NewRequest(method, url, body)
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