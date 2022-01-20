package data

import (
	"SistemVaksinAPI/features/requestvaksin"

	"gorm.io/gorm"
)

type Requestvaksin struct {
	gorm.Model
	ID            int
	Nama          string
	NIK           int
	JenisKelamin  string
	TanggalLahir  string
	Nomor         string
	Alamat        string
	Provinsi      string
	Kota          string
	Kecamatan     string
	Kelurahan     string
	UserID        int
	VaksinID_satu int
	Status_satu   string
	VaksinID_dua  int
	Status_dua    string
}

func toCore(r *Requestvaksin) requestvaksin.RequestvaksinCore {
	return requestvaksin.RequestvaksinCore{
		ID:            int(r.ID),
		Nama:          r.Nama,
		NIK:           r.NIK,
		JenisKelamin:  r.JenisKelamin,
		TanggalLahir:  r.TanggalLahir,
		Nomor:         r.Nomor,
		Alamat:        r.Alamat,
		Provinsi:      r.Provinsi,
		Kota:          r.Kota,
		Kecamatan:     r.Kecamatan,
		Kelurahan:     r.Kelurahan,
		UserID:        r.UserID,
		VaksinID_satu: r.VaksinID_satu,
		Status_satu:   r.Status_satu,
		VaksinID_dua:  r.VaksinID_dua,
		Status_dua:    r.Status_dua,
	}
}

func fromCore(core requestvaksin.RequestvaksinCore) Requestvaksin {
	return Requestvaksin{
		ID:            int(core.ID),
		Nama:          core.Nama,
		NIK:           core.NIK,
		JenisKelamin:  core.JenisKelamin,
		TanggalLahir:  core.TanggalLahir,
		Nomor:         core.Nomor,
		Alamat:        core.Alamat,
		Provinsi:      core.Provinsi,
		Kota:          core.Kota,
		Kecamatan:     core.Kecamatan,
		Kelurahan:     core.Kelurahan,
		UserID:        core.UserID,
		VaksinID_satu: core.VaksinID_satu,
		Status_satu:   core.Status_satu,
		VaksinID_dua:  core.VaksinID_dua,
		Status_dua:    core.Status_dua,
	}
}

func toList(resp []Requestvaksin) []requestvaksin.RequestvaksinCore {
	r := []requestvaksin.RequestvaksinCore{}

	for _, value := range resp {
		r = append(r, toCore(&value))
	}

	return r
}
