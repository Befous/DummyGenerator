package main

import (
	"fmt"

	"github.com/Befous/DummyGenerator/controllers"
	"github.com/Befous/DummyGenerator/utils"
)

func main() {
	err := utils.LoadLokasiFromJSON("data_alamat.json")
	if err != nil {
		fmt.Println("Error loading lokasi data:", err)
		return
	}
	alamat := controllers.GenerateRandomAlamat()
	fmt.Print(alamat)
}
