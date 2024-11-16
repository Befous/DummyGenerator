package utils

import (
	"fmt"
	"math/rand"

	"github.com/Befous/DummyGenerator/models"
)

func GetRandomProvinsi(r *rand.Rand) string {
	provinsi := []string{"Jambi", "Bangka Belitung"}
	return provinsi[r.Intn(len(provinsi))]
}

func GetRandomKotaKabupaten(r *rand.Rand, provinsi string) models.DataAlamat {
	var filteredLokasi []models.DataAlamat
	for _, l := range lokasi {
		if l.Provinsi == provinsi {
			filteredLokasi = append(filteredLokasi, l)
		}
	}
	return filteredLokasi[r.Intn(len(filteredLokasi))]
}

func GetRandomKecamatan(r *rand.Rand, kota models.DataAlamat) models.DataAlamat {
	var filteredLokasi []models.DataAlamat
	for _, l := range lokasi {
		if l.Kota_Kabupaten == kota.Kota_Kabupaten {
			filteredLokasi = append(filteredLokasi, l)
		}
	}
	return filteredLokasi[r.Intn(len(filteredLokasi))]
}

func GetRandomDesaKelurahan(r *rand.Rand, kecamatan models.DataAlamat) models.DataAlamat {
	var filteredLokasi []models.DataAlamat
	for _, l := range lokasi {
		if l.Kecamatan == kecamatan.Kecamatan {
			filteredLokasi = append(filteredLokasi, l)
		}
	}
	return filteredLokasi[r.Intn(len(filteredLokasi))]
}

func GetRandomKodePos(r *rand.Rand, desa models.DataAlamat) models.DataAlamat {
	var filteredLokasi []models.DataAlamat
	for _, l := range lokasi {
		if l.Desa_Kelurahan == desa.Desa_Kelurahan {
			filteredLokasi = append(filteredLokasi, l)
		}
	}
	return filteredLokasi[r.Intn(len(filteredLokasi))]
}

func GenerateRandomAlamatLengkap(r *rand.Rand, alamat models.DataAlamat) string {
	namaJalan := "Jl Caringin Gg Lumbung No 10"
	rt := rand.Intn(99) + 1
	rw := rand.Intn(99) + 1

	rtFormatted := fmt.Sprintf("%02d", rt)
	rwFormatted := fmt.Sprintf("%02d", rw)
	alamatLengkap := namaJalan + " RT " + rtFormatted + " RW " + rwFormatted + ", " + alamat.Desa_Kelurahan + ", " + alamat.Kecamatan + ", " + alamat.Kota_Kabupaten + ", " + alamat.Provinsi + " - " + alamat.Kode_Pos

	return alamatLengkap
}
