package main

import (
	"fmt"

	generator "github.com/Befous/DummyGenerator"
)

func main() {
	alamat := generator.GenerateRandomAlamat()
	fmt.Println("Data Alamat: " + fmt.Sprintf("%+v", alamat))
}
