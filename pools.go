package proxmox

type Pool struct {
	PoolID  string `json:"poolid"`
	Comment string `json:"comment"`
}

type PoolList []*Pool

func GetPools() (PoolList, error) {
	result := &PoolListResponse{}
	request := client.R()
	request.SetResult(result)
	_, err := request.Get("/pools")
	if err != nil {
		return PoolList{}, err
	}
	return result.Data, nil
}
