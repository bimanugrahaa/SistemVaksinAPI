package data

import (
	"SistemVaksinAPI/features/vaksin"

	"gorm.io/gorm"
)

type Vaksin struct {
	gorm.Model
	ID          int
	Jenisvaksin string
	Jadwal      string
	Waktu       string
	Stokvaksin  int
}

func toCore(v *Vaksin) vaksin.VaksinCore {
	return vaksin.VaksinCore{
		ID:          int(v.ID),
		Jenisvaksin: v.Jenisvaksin,
		Jadwal:      v.Jadwal,
		Waktu:       v.Waktu,
		Stokvaksin:  v.Stokvaksin,
	}
}

func fromCore(core vaksin.VaksinCore) Vaksin {
	return Vaksin{
		ID:          int(core.ID),
		Jenisvaksin: core.Jenisvaksin,
		Jadwal:      core.Jadwal,
		Waktu:       core.Waktu,
		Stokvaksin:  core.Stokvaksin,
	}
}
