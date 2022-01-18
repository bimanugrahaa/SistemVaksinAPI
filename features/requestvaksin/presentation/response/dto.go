package response

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

func FromCore(core requestvaksin.RequestvaksinCore) Requestvaksin {
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

func FromCoreSlice(core []requestvaksin.RequestvaksinCore) []Requestvaksin {
	var requestvaksinArray []Requestvaksin
	for key := range core {
		requestvaksinArray = append(requestvaksinArray, FromCore(core[key]))
	}

	return requestvaksinArray
}
