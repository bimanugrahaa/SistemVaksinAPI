package response

import "SistemVaksinAPI/features/faskes"

type Faskes struct {
	ID        int      `json:"id"`
	Nama      string   `json:"nama"`
	Alamat    string   `json:"alamat"`
	Provinsi  string   `json:"provinsi"`
	Kota      string   `json:"kota"`
	Kecamatan string   `json:"kecamatan"`
	Kelurahan string   `json:"kelurahan"`
	Vaksin    []Vaksin `json:"vaksin"`
}

type Vaksin struct {
	ID          int    `json:"vaksin_id"`
	Jenisvaksin string `json:"jenisvaksin"`
	Jadwal      string `json:"jadwal"`
	Waktu       string `json:"waktu"`
	Stokvaksin  int    `json:"stokvaksin"`
	FaskesID    int    `json:"faskes_id"`
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
		Vaksin:    FromVaksinSlice(core.Vaksin),
	}
}

func FromCoreSlice(core []faskes.FaskesCore) []Faskes {
	var faskesArray []Faskes
	for key := range core {
		faskesArray = append(faskesArray, FromCore(core[key]))
	}

	return faskesArray
}

func FromVaksinCore(vc faskes.VaksinCore) Vaksin {
	return Vaksin{
		ID:          vc.ID,
		Jenisvaksin: vc.Jenisvaksin,
		Jadwal:      vc.Jadwal,
		Waktu:       vc.Waktu,
		Stokvaksin:  vc.Stokvaksin,
		FaskesID:    vc.FaskesID,
	}
}

func FromVaksinSlice(vc []faskes.VaksinCore) []Vaksin {
	var vaksinArray []Vaksin
	for key := range vc {
		vaksinArray = append(vaksinArray, FromVaksinCore(vc[key]))
	}

	return vaksinArray
}
