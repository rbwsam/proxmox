package main

import (
	"github.com/rbwsam/proxmox"
	"os"
)

func main() {
	proxmox.Configure(os.Getenv("HOST"), true)
	err := proxmox.Login("root@pam", os.Getenv("PASSWORD"))

	if err != nil {
		panic(err)
	}

	_, err = proxmox.GetStorage("nfs")

	if err != nil {
		panic(err)
	}
}
