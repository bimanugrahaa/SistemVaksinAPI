package request

import (
	"SistemVaksinAPI/features/faskes"
	"SistemVaksinAPI/features/vaksin"
)

type Faskes struct {
	ID        int        `json:"id"`
	Nama      string     `json:"nama"`
	Alamat    string     `json:"alamat"`
	Provinsi  string     `json:"provinsi"`
	Kota      string     `json:"kota"`
	Kecamatan string     `json:"kecamatan"`
	Kelurahan string     `json:"kelurahan"`
	VaksinID  int        `json:"vaksin_id"`
	Vaksin    VaksinDesc `json:"vaksin"`
}

type VaksinDesc struct {
	ID          int    `json:"id"`
	Jenisvaksin string `json:"jenisvaksin"`
	Jadwal      string `json:"jadwal"`
	Waktu       string `json:"waktu"`
	Stokvaksin  int    `json:"stokvaksin"`
}

func ToCore(f Faskes) faskes.FaskesCore {
	return faskes.FaskesCore{
		ID:        int(f.ID),
		Nama:      f.Nama,
		Alamat:    f.Alamat,
		Provinsi:  f.Provinsi,
		Kota:      f.Kota,
		Kecamatan: f.Kecamatan,
		Kelurahan: f.Kelurahan,
		// Vaksin:    ToVaksinCore(req.Jenisvaksin),
	}
}

func ToVaksinCore(req VaksinDesc) vaksin.VaksinCore {
	return vaksin.VaksinCore{
		Jenisvaksin: req.Jenisvaksin,
		Jadwal:      req.Jadwal,
		Waktu:       req.Waktu,
		Stokvaksin:  req.Stokvaksin,
	}
}
