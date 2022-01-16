package response

import "SistemVaksinAPI/features/faskes"

type Faskes struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Provinsi  string `json:"provinsi"`
	Kota      string `json:"kota"`
	Kecamatan string `json:"kecamatan"`
	Kelurahan string `json:"kelurahan"`
}

func FromCore(core faskes.FaskesCore) Faskes {
	return Faskes{
		ID:        int(core.ID),
		Nama:      core.Nama,
		Alamat:    core.Alamat,
		Provinsi:  core.Provinsi,
		Kota:      core.Kota,
		Kecamatan: core.Kecamatan,
		Kelurahan: core.Kelurahan,
	}
}

func FromCoreSlice(core []faskes.FaskesCore) []Faskes {
	var faskesArray []Faskes
	for key := range core {
		faskesArray = append(faskesArray, FromCore(core[key]))
	}

	return faskesArray
}
