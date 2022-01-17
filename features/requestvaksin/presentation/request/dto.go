package request

import "SistemVaksinAPI/features/requestvaksin"

type Requestvaksin struct {
	ID              int    `json:"id"`
	Jenisvaksin     string `json:"jenisvaksin"`
	Dosis_pertama   bool   `json:"dosis_pertama"`
	Dosis_kedua     bool   `json:"dosis_kedua"`
	Status_satu     string `json:"status_satu"`
	Status_dua      string `json:"status_dua"`
	Nama            string `json:"nama"`
	NIK             int    `json:"nik"`
	Jam_pertama     string `json:"jam_pertama"`
	Jam_Kedua       string `json:"jam_kedua"`
	Tanggal_pertama string `json:"tanggal_pertama"`
	Tanggal_kedua   string `json:"tanggal_kedua"`
	Faskes          string `json:"faskes"`
	Provinsi        string `json:"provinsi"`
	Kota            string `json:"kota"`
	Kelurahan       string `json:"kelurahan"`
	Kecamatan       string `json:"kecamatan"`
}

func ToCore(r *Requestvaksin) requestvaksin.RequestvaksinCore {
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
