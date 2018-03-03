package main

import (
	"github.com/rbwsam/proxmox"
	"os"
	"fmt"
)

func main() {
	proxmox.Configure(os.Getenv("HOST"), true)

	err := proxmox.Login("root@pam", os.Getenv("PASSWORD"))
	if err != nil {
		panic(err)
	}

	storageList, err := proxmox.GetStorageList()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", storageList[0])

	storageList, err = proxmox.GetStorageByType("dir")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", storageList[0])

	storage, err := proxmox.GetStorage("local-lvm")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", storage)

	pools, err := proxmox.GetPools()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", pools[0])
}
