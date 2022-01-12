package response

import "SistemVaksinAPI/features/faskes"

type Faskes struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Provinsi  string `json:"provinsi"`
	Kecamatan string `json:"kecamatan"`
	Kelurahan string `json:"kelurahan"`
}

func FromCore(core faskes.FaskesCore) Faskes {
	return Faskes{
		ID:        int(core.ID),
		Nama:      core.Nama,
		Alamat:    core.Alamat,
		Provinsi:  core.Provinsi,
		Kecamatan: core.Kecamatan,
		Kelurahan: core.Kelurahan,
	}
}
