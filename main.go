package main

import (
	"fmt"

	"github.com/Befous/DummyGenerator/controllers"
)

func main() {
	alamat := controllers.GenerateRandomAlamat()
	fmt.Print(alamat)
}
