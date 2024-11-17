package test

import (
	"fmt"
	"testing"

	generator "github.com/Befous/DummyGenerator"
	"github.com/Befous/DummyGenerator/models"
)

func TestGetRandomProvinsi(t *testing.T) {
	provinsi := generator.GetRandomProvinsi()
	fmt.Println("Provinsi: " + provinsi)
}
func TestGetRandomKotaKabupaten(t *testing.T) {
	provinsi := generator.GetRandomProvinsi()
	kotaKabupaten := generator.GetRandomKotaKabupaten(provinsi)
	alamat := models.DataAlamat{
		Kota_Kabupaten: kotaKabupaten.Kota_Kabupaten,
		Ibu_Kota:       kotaKabupaten.Ibu_Kota,
		Kode_Kota:      kotaKabupaten.Kode_Kota,
		Kode_Bandara:   kotaKabupaten.Kode_Bandara,
		Provinsi:       provinsi,
	}
	fmt.Println("Provinsi: " + provinsi)
	fmt.Println("Kota/Kabupaten: " + kotaKabupaten.Kota_Kabupaten)
	fmt.Println("Data Alamat: " + fmt.Sprintf("%+v", alamat))
}
func TestGetRandomKecamatan(t *testing.T) {
	provinsi := generator.GetRandomProvinsi()
	kotaKabupaten := generator.GetRandomKotaKabupaten(provinsi)
	kecamatan := generator.GetRandomKecamatan(kotaKabupaten)
	alamat := models.DataAlamat{
		Kecamatan:      kecamatan.Kecamatan,
		Kota_Kabupaten: kotaKabupaten.Kota_Kabupaten,
		Ibu_Kota:       kecamatan.Ibu_Kota,
		Kode_Kota:      kecamatan.Kode_Kota,
		Kode_Bandara:   kecamatan.Kode_Bandara,
		Provinsi:       provinsi,
	}
	fmt.Println("Provinsi: " + provinsi)
	fmt.Println("Kota/Kabupaten: " + kotaKabupaten.Kota_Kabupaten)
	fmt.Println("Kecamatan: " + kecamatan.Kecamatan)
	fmt.Println("Data Alamat: " + fmt.Sprintf("%+v", alamat))
}
func TestGetRandomDesaKelurahan(t *testing.T) {
	provinsi := generator.GetRandomProvinsi()
	kotaKabupaten := generator.GetRandomKotaKabupaten(provinsi)
	kecamatan := generator.GetRandomKecamatan(kotaKabupaten)
	desaKelurahan := generator.GetRandomDesaKelurahan(kecamatan)
	alamat := models.DataAlamat{
		Desa_Kelurahan: desaKelurahan.Desa_Kelurahan,
		Kecamatan:      kecamatan.Kecamatan,
		Kota_Kabupaten: kotaKabupaten.Kota_Kabupaten,
		Ibu_Kota:       desaKelurahan.Ibu_Kota,
		Kode_Kota:      desaKelurahan.Kode_Kota,
		Kode_Bandara:   desaKelurahan.Kode_Bandara,
		Provinsi:       provinsi,
	}
	fmt.Println("Provinsi: " + provinsi)
	fmt.Println("Kota/Kabupaten: " + kotaKabupaten.Kota_Kabupaten)
	fmt.Println("Kecamatan: " + kecamatan.Kecamatan)
	fmt.Println("Desa/Kelurahan: " + desaKelurahan.Desa_Kelurahan)
	fmt.Println("Alamat_Lengkap: " + alamat.Alamat_Lengkap)
	fmt.Println("Data Alamat: " + fmt.Sprintf("%+v", alamat))
}
func TestGetRandomKodePos(t *testing.T) {
	provinsi := generator.GetRandomProvinsi()
	kotaKabupaten := generator.GetRandomKotaKabupaten(provinsi)
	kecamatan := generator.GetRandomKecamatan(kotaKabupaten)
	desaKelurahan := generator.GetRandomDesaKelurahan(kecamatan)
	kodePos := generator.GetRandomKodePos(desaKelurahan)
	alamat := models.DataAlamat{
		Kode_Pos:       kodePos.Kode_Pos,
		Desa_Kelurahan: desaKelurahan.Desa_Kelurahan,
		Kecamatan:      kecamatan.Kecamatan,
		Kota_Kabupaten: kotaKabupaten.Kota_Kabupaten,
		Ibu_Kota:       kodePos.Ibu_Kota,
		Kode_Kota:      kodePos.Kode_Kota,
		Kode_Bandara:   kodePos.Kode_Bandara,
		Provinsi:       provinsi,
	}
	fmt.Println("Provinsi: " + provinsi)
	fmt.Println("Kota/Kabupaten: " + kotaKabupaten.Kota_Kabupaten)
	fmt.Println("Kecamatan: " + kecamatan.Kecamatan)
	fmt.Println("Desa/Kelurahan: " + desaKelurahan.Desa_Kelurahan)
	fmt.Println("Kode Pos: " + fmt.Sprintf("%d", kodePos.Kode_Pos))
	fmt.Println("Data Alamat: " + fmt.Sprintf("%+v", alamat))
}
func TestGenerateRandomAlamatLengkap(t *testing.T) {
	provinsi := generator.GetRandomProvinsi()
	kotaKabupaten := generator.GetRandomKotaKabupaten(provinsi)
	kecamatan := generator.GetRandomKecamatan(kotaKabupaten)
	desaKelurahan := generator.GetRandomDesaKelurahan(kecamatan)
	kodePos := generator.GetRandomKodePos(desaKelurahan)
	alamat := models.DataAlamat{
		Kode_Pos:       kodePos.Kode_Pos,
		Desa_Kelurahan: desaKelurahan.Desa_Kelurahan,
		Kecamatan:      kecamatan.Kecamatan,
		Kota_Kabupaten: kotaKabupaten.Kota_Kabupaten,
		Ibu_Kota:       kodePos.Ibu_Kota,
		Kode_Kota:      kodePos.Kode_Kota,
		Kode_Bandara:   kodePos.Kode_Bandara,
		Provinsi:       provinsi,
	}
	alamat.Alamat_Lengkap = generator.GenerateRandomAlamatLengkap(alamat)
	fmt.Println("Provinsi: " + provinsi)
	fmt.Println("Kota/Kabupaten: " + kotaKabupaten.Kota_Kabupaten)
	fmt.Println("Kecamatan: " + kecamatan.Kecamatan)
	fmt.Println("Desa/Kelurahan: " + desaKelurahan.Desa_Kelurahan)
	fmt.Println("Kode Pos: " + fmt.Sprintf("%d", kodePos.Kode_Pos))
	fmt.Println("Alamat_Lengkap: " + alamat.Alamat_Lengkap)
	fmt.Println("Data Alamat: " + fmt.Sprintf("%+v", alamat))
}
func TestGenerateRandomAlamat(t *testing.T) {
	alamat := generator.GenerateRandomAlamat()
	fmt.Println("Data Alamat: " + fmt.Sprintf("%+v", alamat))
}
