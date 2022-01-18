package vaksin

type VaksinCore struct {
	ID          int
	Jenisvaksin string
	Jadwal      string
	Waktu       string
	Stokvaksin  int
	FaskesID    int
}

type Bussiness interface {
	CreateVaksin(data VaksinCore) (resp VaksinCore, err error)
	GetVaksinByFaskesID(id int) (resp []VaksinCore, err error)
	GetVaksinByID(id int) (resp VaksinCore, err error)
	UpdateVaksinByID(data VaksinCore) (resp VaksinCore, err error)
}

type Data interface {
	InsertVaksin(data VaksinCore) (resp VaksinCore, err error)
	SelectVaksinByFaskesID(id int) (resp []VaksinCore, err error)
	SelectVaksinByID(id int) (resp VaksinCore, err error)
	EditVaksinByID(data VaksinCore) (resp VaksinCore, err error)
}
