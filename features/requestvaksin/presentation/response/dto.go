package response

import "SistemVaksinAPI/features/requestvaksin"

type Requestvaksin struct {
	ID           int    `json:"id"`
	Nama         string `json:"nama"`
	NIK          int    `json:"nik"`
	JenisKelamin string `json:"jeniskelamin"`
	TanggalLahir string `json:"tanggallahir"`
	Nomor        string `json:"nomor"`
	Alamat       string `json:"alamat"`

	Provinsi  string `json:"provinsi"`
	Kota      string `json:"kota"`
	Kecamatan string `json:"kecamatan"`
	Kelurahan string `json:"kelurahan"`

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

func FromCore(core requestvaksin.RequestvaksinCore) Requestvaksin {
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
		Vaksin_satu:   FromVaksinCore(core.Vaksin_satu),
		VaksinID_dua:  core.VaksinID_dua,
		Status_dua:    core.Status_dua,
		Vaksin_dua:    FromVaksinCore(core.Vaksin_dua),
	}
}

func FromVaksinCore(vc requestvaksin.VaksinCore) Vaksin {
	return Vaksin{
		ID:          vc.ID,
		Jenisvaksin: vc.Jenisvaksin,
		Jadwal:      vc.Jadwal,
		Waktu:       vc.Waktu,
		FaskesID:    vc.FaskesID,
		Faskes:      FromFaskesCore(vc.Faskes),
	}
}

func FromFaskesCore(vc requestvaksin.FaskesCore) Faskes {
	return Faskes{
		ID:   vc.ID,
		Nama: vc.Nama,
	}
}

func FromCoreSlice(core []requestvaksin.RequestvaksinCore) []Requestvaksin {
	var requestvaksinArray []Requestvaksin
	for key := range core {
		requestvaksinArray = append(requestvaksinArray, FromCore(core[key]))
	}

	return requestvaksinArray
}

func FromCoreLogin(core requestvaksin.RequestvaksinCore) Requestvaksin {
	return Requestvaksin{
		ID:   core.ID,
		Nama: core.Nama,
		NIK:  core.NIK,
	}
}
