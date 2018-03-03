package proxmox

type LoginResponse struct {
	Data struct {
		Username            string                    `json:"username"`
		Capabilities        map[string]map[string]int `json:"cap"`
		CSRFPreventionToken string                    `json:"CSRFPreventionToken"`
		Ticket              string                    `json:"ticket"`
	} `json:"data"`
}

type NodeListResponse struct {
	Data NodeList `json:"Data"`
}

type PoolListResponse struct {
	Data PoolList `json:"data"`
}

type StorageListResponse struct {
	Data StorageList `json:"data"`
}

type StorageResponse struct {
	Data *Storage `json:"data"`
}

type VirtualMachineListResponse struct {
	Data VirtualMachineList `json:"Data"`
}
