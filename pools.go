package proxmox

import "github.com/rbwsam/proxmox/models"

func GetPools() (models.PoolList, error) {
	result := &models.PoolListResponse{}
	request := client.R()
	request.SetResult(result)
	_, err := request.Get("/pools")
	if err != nil {
		return models.PoolList{}, err
	}
	return result.Data, nil
}
