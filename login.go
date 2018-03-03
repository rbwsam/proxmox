package proxmox

import (
	"net/http"
)

func Login(username, password string) error {
	data := map[string]string{
		"username": username,
		"password": password,
	}
	result := &LoginResponse{}

	request := client.R()
	request.SetFormData(data)
	request.SetResult(result)
	_, err := request.Post("/access/ticket")
	if err != nil {
		return err
	}
	client.SetCookie(&http.Cookie{
		Name:  "PVEAuthCookie",
		Value: result.Data.Ticket,
	})
	return nil
}
