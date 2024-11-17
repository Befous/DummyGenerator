package models

type DataAlamat struct {
	Kode_Pos       int    `json:"kode_pos,omitempty" bson:"kode_pos,omitempty"`
	Desa_Kelurahan string `json:"desa_kelurahan,omitempty" bson:"desa_kelurahan,omitempty"`
	Kecamatan      string `json:"kecamatan,omitempty" bson:"kecamatan,omitempty"`
	Kota_Kabupaten string `json:"kota_kabupaten,omitempty" bson:"kota_kabupaten,omitempty"`
	Ibu_Kota       string `json:"ibu_kota,omitempty" bson:"ibu_kota,omitempty"`
	Kode_Kota      string `json:"kode_kota,omitempty" bson:"kode_kota,omitempty"`
	Kode_Bandara   string `json:"kode_bandara,omitempty" bson:"kode_bandara,omitempty"`
	Provinsi       string `json:"provinsi,omitempty" bson:"provinsi,omitempty"`
	Alamat_Lengkap string `json:"alamat_lengkap,omitempty" bson:"alamat_lengkap,omitempty"`
}
