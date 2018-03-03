package proxmox

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/resty.v1"
)

var client *resty.Client

func Configure(fqdn string, allowInsecure bool) {
	hostURL := fmt.Sprintf("https://%s:8006/api2/json", fqdn)
	client = resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: allowInsecure})
	client.SetHostURL(hostURL)
	client.SetRESTMode()
	client.Debug = true
}
