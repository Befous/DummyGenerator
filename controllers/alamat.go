package controllers

import (
	"math/rand"
	"time"

	"github.com/Befous/DummyGenerator/models"
	"github.com/Befous/DummyGenerator/utils"
)

func GenerateRandomAlamat() models.DataAlamat {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	provinsi := utils.GetRandomProvinsi(r)
	kotaKabupaten := utils.GetRandomKotaKabupaten(r, provinsi)
	kecamatan := utils.GetRandomKecamatan(r, kotaKabupaten)
	desaKelurahan := utils.GetRandomDesaKelurahan(r, kecamatan)
	kodePos := utils.GetRandomKodePos(r, desaKelurahan)
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
	alamat.Alamat_Lengkap = utils.GenerateRandomAlamatLengkap(r, alamat)

	return alamat
}
