package data

import (
	"SistemVaksinAPI/features/requestvaksin"

	"gorm.io/gorm"
)

type Requestvaksin struct {
	gorm.Model
	ID              int
	Jenisvaksin     string
	Dosis_pertama   bool
	Dosis_kedua     bool
	Status_satu     string
	Status_dua      string
	Nama            string
	NIK             int
	Jam_pertama     string
	Jam_Kedua       string
	Tanggal_pertama string
	Tanggal_kedua   string
	Faskes          string
	Provinsi        string
	Kota            string
	Kelurahan       string
	Kecamatan       string
}

func toCore(r *Requestvaksin) requestvaksin.RequestvaksinCore {
	return requestvaksin.RequestvaksinCore{
		ID:              int(r.ID),
		Jenisvaksin:     r.Jenisvaksin,
		Dosis_pertama:   r.Dosis_pertama,
		Dosis_kedua:     r.Dosis_kedua,
		Status_satu:     r.Status_satu,
		Status_dua:      r.Status_dua,
		Nama:            r.Nama,
		NIK:             r.NIK,
		Jam_pertama:     r.Jam_pertama,
		Jam_Kedua:       r.Jam_Kedua,
		Tanggal_pertama: r.Tanggal_pertama,
		Tanggal_kedua:   r.Tanggal_kedua,
		Faskes:          r.Faskes,
		Provinsi:        r.Provinsi,
		Kota:            r.Kota,
		Kelurahan:       r.Kelurahan,
		Kecamatan:       r.Kecamatan,
	}
}

func fromCore(core requestvaksin.RequestvaksinCore) Requestvaksin {
	return Requestvaksin{
		ID:              int(core.ID),
		Jenisvaksin:     core.Jenisvaksin,
		Dosis_pertama:   core.Dosis_pertama,
		Dosis_kedua:     core.Dosis_kedua,
		Status_satu:     core.Status_satu,
		Status_dua:      core.Status_dua,
		Nama:            core.Nama,
		NIK:             core.NIK,
		Jam_pertama:     core.Jam_pertama,
		Jam_Kedua:       core.Jam_Kedua,
		Tanggal_pertama: core.Tanggal_pertama,
		Tanggal_kedua:   core.Tanggal_kedua,
		Faskes:          core.Faskes,
		Provinsi:        core.Provinsi,
		Kota:            core.Kota,
		Kelurahan:       core.Kelurahan,
		Kecamatan:       core.Kecamatan,
	}
}

func toList(resp []Requestvaksin) []requestvaksin.RequestvaksinCore {
	r := []requestvaksin.RequestvaksinCore{}

	for _, value := range resp {
		r = append(r, toCore(&value))
	}

	return r
}
