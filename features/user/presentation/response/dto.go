package response

import "SistemVaksinAPI/features/user"

type User struct {
	UserID        int             `json:"id"`
	Password      string          `json:"password"`
	Email         string          `json:"email"`
	Namalengkap   string          `json:"namalengkap"`
	NIK           int             `json:"nik"`
	Token         string          `json:"token"`
	Requestvaksin []Requestvaksin `json:"requestvaksin"`
}

type Requestvaksin struct {
	ID            int    `json:"id"`
	Nama          string `json:"nama"`
	NIK           int    `json:"nik"`
	JenisKelamin  string `json:"jeniskelamin"`
	TanggalLahir  string `json:"tanggallahir"`
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

func FromCore(core user.UserCore) User {
	return User{
		UserID:        core.UserID,
		Namalengkap:   core.Namalengkap,
		NIK:           core.NIK,
		Email:         core.Email,
		Token:         core.Token,
		Requestvaksin: FromRequestVaksinSlice(core.RequestVaksin),
	}
}

func FromCoreLogin(core user.UserCore) User {
	return User{
		UserID:      core.UserID,
		Namalengkap: core.Namalengkap,
		NIK:         core.NIK,
		Email:       core.Email,
		Token:       core.Token,
	}
}

func FromRequestVaksinSlice(vc []user.RequestvaksinCore) []Requestvaksin {
	var requestVaksinArray []Requestvaksin

	for key := range vc {
		requestVaksinArray = append(requestVaksinArray, FromRequestVaksinCore(vc[key]))
	}

	return requestVaksinArray
}

func FromRequestVaksinCore(r user.RequestvaksinCore) Requestvaksin {
	return Requestvaksin{
		ID:            int(r.ID),
		Nama:          r.Nama,
		NIK:           r.NIK,
		JenisKelamin:  r.JenisKelamin,
		TanggalLahir:  r.TanggalLahir,
		UserID:        r.UserID,
		VaksinID_satu: r.VaksinID_satu,
		Status_satu:   r.Status_satu,
		Vaksin_satu:   FromVaksinCore(r.Vaksin_satu),
		VaksinID_dua:  r.VaksinID_dua,
		Status_dua:    r.Status_dua,
		Vaksin_dua:    FromVaksinCore(r.Vaksin_dua),
	}
}

func FromVaksinCore(req user.VaksinCore) Vaksin {
	return Vaksin{
		ID:          req.ID,
		Jenisvaksin: req.Jenisvaksin,
		Jadwal:      req.Jadwal,
		Waktu:       req.Waktu,
		FaskesID:    req.FaskesID,
		Faskes:      FromFaskesCore(req.Faskes),
	}
}

func FromFaskesCore(req user.FaskesCore) Faskes {
	return Faskes{
		ID:   req.ID,
		Nama: req.Nama,
	}
}
