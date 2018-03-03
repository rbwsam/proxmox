package proxmox

import (
	"github.com/rbwsam/proxmox/models"
)

func GetStorage(storage string) (*models.Storage, error) {
	pathParams := make(map[string]string)
	pathParams["storage"] = storage
	result := &models.StorageResponse{}
	request := client.R()
	request.SetResult(result)
	request.SetPathParams(pathParams)
	_, err := request.Get("/storage/{storage}")
	if err != nil {
		return &models.Storage{}, err
	}
	return result.Data, nil
}

func GetStorageList() (models.StorageList, error) {
	result := &models.StorageListResponse{}
	request := client.R()
	request.SetResult(result)
	_, err := request.Get("/storage")
	if err != nil {
		return models.StorageList{}, err
	}
	return result.Data, nil
}

func GetStorageByType(storageType string) (models.StorageList, error) {
	result := &models.StorageListResponse{}
	request := client.R()
	request.SetQueryParam("type", storageType)
	request.SetResult(result)
	_, err := request.Get("/storage")
	if err != nil {
		return models.StorageList{}, err
	}
	return result.Data, nil
}