package proxmox

import (
	"github.com/rbwsam/proxmox/models"
)

func Login(username, password string) error {
	data := map[string]string{
		"username": username,
		"password": password,
	}
	body := generateRequestBody(data)

	resp, err := Post("/access/ticket", body)
	if err != nil {
		return err
	}

	result := &models.LoginResponse{}

	err = unmarshalResponse(resp, result)
	if err != nil {
		return err
	}

	setCookie(result.Data.Ticket)
	return nil
}
