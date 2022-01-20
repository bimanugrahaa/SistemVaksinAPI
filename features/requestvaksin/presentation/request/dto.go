package request

import (
	"SistemVaksinAPI/features/faskes"
	"SistemVaksinAPI/features/requestvaksin"
	"SistemVaksinAPI/features/vaksin"
)

type Requestvaksin struct {
	ID            int    `json:"id"`
	Nama          string `json:"nama"`
	NIK           int    `json:"nik"`
	JenisKelamin  string `json:"jeniskelamin"`
	TanggalLahir  string `json:"tanggallahir"`
	Nomor         string `json:"nomor"`
	Alamat        string `json:"alamat"`
	Provinsi      string `json:"provinsi"`
	Kota          string `json:"kota"`
	Kecamatan     string `json:"kecamatan"`
	Kelurahan     string `json:"kelurahan"`
	UserID        int    `json:"user_id"`
	VaksinID_satu int    `json:"vaksinID_satu"`
	Status_satu   string `json:"status_satu"`
	Vaksin_satu   Vaksin `json:"vaksin_satu"`
	VaksinID_dua  int    `json:"vaksinID_dua"`
	Status_dua    string `json:"status_dua"`
	Vaksin_dua    Vaksin `json:"vaksin_dua"`
}

type Vaksin struct {
	ID          int    `json:"id"`
	Jenisvaksin string `json:"jenisvaksin"`
	Jadwal      string `json:"jadwal"`
	Waktu       string `json:"waktu"`
	FaskesID    int    `json:"faskes_id"`
	Faskes      Faskes `json:"faskes"`
}

type Faskes struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

func ToCore(r Requestvaksin) requestvaksin.RequestvaksinCore {
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

func ToVaksinCore(req Vaksin) vaksin.VaksinCore {
	return vaksin.VaksinCore{
		Jenisvaksin: req.Jenisvaksin,
		Jadwal:      req.Jadwal,
		Waktu:       req.Waktu,
		FaskesID:    req.FaskesID,
	}
}

func ToVFaskesCore(req Faskes) faskes.FaskesCore {
	return faskes.FaskesCore{
		Nama: req.Nama,
	}
}
