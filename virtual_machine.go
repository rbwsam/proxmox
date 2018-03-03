package proxmox

type VirtualMachine struct {
	DiskRead  int     `json:"diskread"`
	CPUs      int     `json:"cpus"`
	NetIn     int64   `json:"netin"`
	MaxMem    int64   `json:"maxmem"`
	Status    string  `json:"status"`
	Uptime    int     `json:"uptime"`
	DiskWrite int     `json:"diskwrite"`
	Disk      int     `json:"disk"`
	Mem       int64   `json:"mem"`
	MaxDisk   int64   `json:"maxdisk"`
	NetOut    int64   `json:"netout"`
	PID       string  `json:"pid"`
	CPU       float64 `json:"cpu"`
	Name      string  `json:"name"`
	Template  string  `json:"template"`
	VMID      int     `json:"vmid"`
}

type VirtualMachineList []*VirtualMachine

func GetVirtualMachines(node string) (VirtualMachineList, error) {
	pathParams := make(map[string]string)
	pathParams["node"] = node
	result := &VirtualMachineListResponse{}
	request := client.R()
	request.SetResult(result)
	request.SetPathParams(pathParams)
	_, err := request.Get("/nodes/{node}/qemu")
	if err != nil {
		return VirtualMachineList{}, err
	}
	return result.Data, nil
}
