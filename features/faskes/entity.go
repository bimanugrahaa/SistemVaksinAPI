package faskes

type FaskesCore struct {
	ID        int
	Vaksin    []VaksinCore
	Nama      string
	Alamat    string
	Provinsi  string
	Kota      string
	Kecamatan string
	Kelurahan string
}

type VaksinCore struct {
	ID          int
	Jenisvaksin string
	Jadwal      string
	Waktu       string
	Stokvaksin  int
	FaskesID    int
}

type Bussiness interface {
	CreateFaskes(data FaskesCore) (resp FaskesCore, err error)
	GetAllFaskes() (resp []FaskesCore)
	GetFaskesByID(id int) (resp FaskesCore, err error)
	GetFaskesByName(data FaskesCore) (resp FaskesCore, err error)
}

type Data interface {
	InsertFaskes(data FaskesCore) (resp FaskesCore, err error)
	SelectAllFaskes() (resp []FaskesCore)
	SelectFaskesByID(id int) (resp FaskesCore, err error)
	SelectFaskesByName(data FaskesCore) (resp FaskesCore, err error)
}
