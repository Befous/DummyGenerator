package generator

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Befous/DummyGenerator/models"
)

func GetRandomProvinsi() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return provinsi[r.Intn(len(provinsi))]
}

func GetRandomKotaKabupaten(provinsi string) models.DataAlamat {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var filteredLokasi []models.DataAlamat
	for _, l := range lokasi {
		if l.Provinsi == provinsi {
			filteredLokasi = append(filteredLokasi, l)
		}
	}
	return filteredLokasi[r.Intn(len(filteredLokasi))]
}

func GetRandomKecamatan(kota models.DataAlamat) models.DataAlamat {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var filteredLokasi []models.DataAlamat
	for _, l := range lokasi {
		if l.Kota_Kabupaten == kota.Kota_Kabupaten {
			filteredLokasi = append(filteredLokasi, l)
		}
	}
	return filteredLokasi[r.Intn(len(filteredLokasi))]
}

func GetRandomDesaKelurahan(kecamatan models.DataAlamat) models.DataAlamat {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var filteredLokasi []models.DataAlamat
	for _, l := range lokasi {
		if l.Kecamatan == kecamatan.Kecamatan {
			filteredLokasi = append(filteredLokasi, l)
		}
	}
	return filteredLokasi[r.Intn(len(filteredLokasi))]
}

func GetRandomKodePos(desa models.DataAlamat) models.DataAlamat {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var filteredLokasi []models.DataAlamat
	for _, l := range lokasi {
		if l.Desa_Kelurahan == desa.Desa_Kelurahan {
			filteredLokasi = append(filteredLokasi, l)
		}
	}
	return filteredLokasi[r.Intn(len(filteredLokasi))]
}

func GenerateRandomAlamatLengkap(alamat models.DataAlamat) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	streetName := namaJalan[r.Intn(len(namaJalan))]
	rt := rand.Intn(99) + 1
	rw := rand.Intn(99) + 1

	rtFormatted := fmt.Sprintf("%02d", rt)
	rwFormatted := fmt.Sprintf("%02d", rw)
	kodePosFormatted := fmt.Sprintf("%d", alamat.Kode_Pos)
	alamatLengkap := streetName + " RT " + rtFormatted + " RW " + rwFormatted + ", " + alamat.Desa_Kelurahan + ", " + alamat.Kecamatan + ", " + alamat.Kota_Kabupaten + ", " + alamat.Provinsi + " - " + kodePosFormatted

	return alamatLengkap
}

func GenerateRandomAlamat() (alamat models.DataAlamat) {
	provinsi := GetRandomProvinsi()
	kotaKabupaten := GetRandomKotaKabupaten(provinsi)
	kecamatan := GetRandomKecamatan(kotaKabupaten)
	desaKelurahan := GetRandomDesaKelurahan(kecamatan)
	kodePos := GetRandomKodePos(desaKelurahan)
	alamat = models.DataAlamat{
		Kode_Pos:       kodePos.Kode_Pos,
		Desa_Kelurahan: desaKelurahan.Desa_Kelurahan,
		Kecamatan:      kecamatan.Kecamatan,
		Kota_Kabupaten: kotaKabupaten.Kota_Kabupaten,
		Ibu_Kota:       kodePos.Ibu_Kota,
		Kode_Kota:      kodePos.Kode_Kota,
		Kode_Bandara:   kodePos.Kode_Bandara,
		Provinsi:       provinsi,
	}
	alamat.Alamat_Lengkap = GenerateRandomAlamatLengkap(alamat)

	return alamat
}
