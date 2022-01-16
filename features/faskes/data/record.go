package data

import (
	"SistemVaksinAPI/features/faskes"

	"gorm.io/gorm"
)

type Faskes struct {
	gorm.Model
	ID        int
	VaksinID  int
	Nama      string
	Alamat    string
	Provinsi  string
	Kota      string
	Kecamatan string
	Kelurahan string
}

func toCore(f *Faskes) faskes.FaskesCore {
	return faskes.FaskesCore{
		ID:        int(f.ID),
		VaksinID:  f.VaksinID,
		Nama:      f.Nama,
		Alamat:    f.Alamat,
		Provinsi:  f.Provinsi,
		Kota:      f.Kota,
		Kecamatan: f.Kecamatan,
		Kelurahan: f.Kelurahan,
	}
}

func toList(resp []Faskes) []faskes.FaskesCore {
	f := []faskes.FaskesCore{}

	for _, value := range resp {
		f = append(f, toCore(&value))
	}

	return f
}

func fromCore(core faskes.FaskesCore) Faskes {
	return Faskes{
		ID:        int(core.ID),
		VaksinID:  core.VaksinID,
		Nama:      core.Nama,
		Alamat:    core.Alamat,
		Provinsi:  core.Provinsi,
		Kota:      core.Kota,
		Kecamatan: core.Kecamatan,
		Kelurahan: core.Kelurahan,
	}
}
