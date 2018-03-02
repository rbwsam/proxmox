package proxmox

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

var server string
var cookie *http.Cookie
var transport *http.Transport
var client *http.Client

func Configure(fqdn string, allowInsecure bool) {
	server = fmt.Sprintf("https://%s:8006/api2/json", fqdn)
	transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: allowInsecure},
	}
	client = &http.Client{
		Transport: transport,
	}
}
