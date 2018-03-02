package models

type LoginResponse struct {
	Data struct {
		Username            string                    `json:"username"`
		Capabilities        map[string]map[string]int `json:"cap"`
		CSRFPreventionToken string                    `json:"CSRFPreventionToken"`
		Ticket              string                    `json:"ticket"`
	} `json:"data"`
}
