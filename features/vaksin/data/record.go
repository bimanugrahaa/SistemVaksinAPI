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
	FaskesID    int
}

func toCore(v *Vaksin) vaksin.VaksinCore {
	return vaksin.VaksinCore{
		ID:          int(v.ID),
		Jenisvaksin: v.Jenisvaksin,
		Jadwal:      v.Jadwal,
		Waktu:       v.Waktu,
		Stokvaksin:  v.Stokvaksin,
		FaskesID:    v.FaskesID,
	}
}

func fromCore(core vaksin.VaksinCore) Vaksin {
	return Vaksin{
		ID:          int(core.ID),
		Jenisvaksin: core.Jenisvaksin,
		Jadwal:      core.Jadwal,
		Waktu:       core.Waktu,
		Stokvaksin:  core.Stokvaksin,
		FaskesID:    core.FaskesID,
	}
}

func toVaksinList(resp []Vaksin) []vaksin.VaksinCore {
	cc := []vaksin.VaksinCore{}

	for _, value := range resp {
		cc = append(cc, toCore(&value))
	}

	return cc
}
