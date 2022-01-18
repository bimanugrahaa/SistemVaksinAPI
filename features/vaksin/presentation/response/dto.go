package response

import "SistemVaksinAPI/features/vaksin"

type Vaksin struct {
	ID          int    `json:"id"`
	Jenisvaksin string `json:"jenisvaksin"`
	Jadwal      string `json:"jadwal"`
	Waktu       string `json:"waktu"`
	Stokvaksin  int    `json:"stokvaksin"`
	FaskesID    int    `json:"faskes_id"`
}

func FromCore(core vaksin.VaksinCore) Vaksin {
	return Vaksin{
		ID:          int(core.ID),
		Jenisvaksin: core.Jenisvaksin,
		Jadwal:      core.Jadwal,
		Waktu:       core.Waktu,
		Stokvaksin:  core.Stokvaksin,
		FaskesID:    core.FaskesID,
	}
}
