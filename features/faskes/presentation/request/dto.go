package request

import (
	"SistemVaksinAPI/features/faskes"
)

type Faskes struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Provinsi  string `json:"provinsi"`
	Kecamatan string `json:"kecamatan"`
	Kelurahan string `json:"kelurahan"`
}

func ToCore(f Faskes) faskes.FaskesCore {
	return faskes.FaskesCore{
		ID:        int(f.ID),
		Nama:      f.Nama,
		Alamat:    f.Alamat,
		Provinsi:  f.Provinsi,
		Kecamatan: f.Kecamatan,
		Kelurahan: f.Kelurahan,
	}
}
