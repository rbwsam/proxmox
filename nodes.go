package proxmox

type Node struct {
	MaxDisk        int64   `json:"maxdisk"`
	Memory         int64   `json:"mem"`
	ID             string  `json:"id"`
	Uptime         int     `json:"uptime"`
	Disk           int     `json:"disk"`
	Type           string  `json:"type"`
	SSLFingerprint string  `json:"ssl_fingerprint"`
	MaxMemory      int64   `json:"maxmem"`
	CPU            float64 `json:"cpu"`
	Level          string  `json:"level"`
	MaxCPU         int     `json:"maxcpu"`
	Status         string  `json:"status"`
	Node           string  `json:"node"`
}

type NodeList []*Node

func GetNodes() (NodeList, error) {
	result := &NodeListResponse{}
	request := client.R()
	request.SetResult(result)
	_, err := request.Get("/nodes")
	if err != nil {
		return NodeList{}, err
	}
	return result.Data, nil
}
