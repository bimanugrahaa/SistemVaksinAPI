package vaksin

type VaksinCore struct {
	ID          int
	Jenisvaksin string
	Jadwal      string
	Waktu       string
	Stokvaksin  int
}

type Bussiness interface {
	CreateVaksin(data VaksinCore) (resp VaksinCore, err error)
	GetVaksinByID(id int) (resp VaksinCore, err error)
}

type Data interface {
	InsertVaksin(data VaksinCore) (resp VaksinCore, err error)
	SelectVaksinByID(id int) (resp VaksinCore, err error)
}
