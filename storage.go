package proxmox

import (
	"github.com/rbwsam/proxmox/models"
)

func GetStorage(storageType string) (*models.StorageList, error) {
	result := &models.GetStorageResponse{}
	request := client.R()
	if storageType != "" {
		request.SetQueryParam("type", storageType)
	}
	request.SetResult(result)
	_, err := request.Get("/storage")
	if err != nil {
		return &models.StorageList{}, err
	}
	return result.Data, nil
}
