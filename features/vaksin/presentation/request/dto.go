package request

import (
	"SistemVaksinAPI/features/vaksin"
)

type Vaksin struct {
	ID          int    `json:"id"`
	Jenisvaksin string `json:"jenisvaksin"`
	Jadwal      string `json:"jadwal"`
	Waktu       string `json:"waktu"`
	Stokvaksin  int    `json:"stokvaksin"`
	FaskesID    int    `json:"faskes_id"`
}

func ToCore(v Vaksin) vaksin.VaksinCore {
	return vaksin.VaksinCore{
		ID:          int(v.ID),
		Jenisvaksin: v.Jenisvaksin,
		Jadwal:      v.Jadwal,
		Waktu:       v.Waktu,
		Stokvaksin:  v.Stokvaksin,
		FaskesID:    v.FaskesID,
	}
}
