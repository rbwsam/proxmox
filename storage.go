package proxmox

type Storage struct {
	// Required
	Storage string `json:"storage"`
	Type    string `json:"type"`
	// Optional
	Content  string `json:"content,omitempty"`
	ThinPool string `json:"thinpool,omitempty"`
	Digest   string `json:"digest,omitempty"`
	VGName   string `json:"vgname,omitempty"`
	Path     string `json:"path,omitempty"`
}

type StorageList []*Storage

func GetStorage(storage string) (*Storage, error) {
	pathParams := make(map[string]string)
	pathParams["storage"] = storage
	result := &StorageResponse{}
	request := client.R()
	request.SetResult(result)
	request.SetPathParams(pathParams)
	_, err := request.Get("/storage/{storage}")
	if err != nil {
		return &Storage{}, err
	}
	return result.Data, nil
}

func GetStorageList() (StorageList, error) {
	result := &StorageListResponse{}
	request := client.R()
	request.SetResult(result)
	_, err := request.Get("/storage")
	if err != nil {
		return StorageList{}, err
	}
	return result.Data, nil
}

func GetStorageByType(storageType string) (StorageList, error) {
	result := &StorageListResponse{}
	request := client.R()
	request.SetQueryParam("type", storageType)
	request.SetResult(result)
	_, err := request.Get("/storage")
	if err != nil {
		return StorageList{}, err
	}
	return result.Data, nil
}
