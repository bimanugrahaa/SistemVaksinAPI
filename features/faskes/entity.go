package faskes

type FaskesCore struct {
	ID        int
	Nama      string
	Alamat    string
	Provinsi  string
	Kecamatan string
	Kelurahan string
}

type Bussiness interface {
	CreateFaskes(data FaskesCore) (resp FaskesCore, err error)
	GetAllFaskes() (resp []FaskesCore)
}

type Data interface {
	InsertFaskes(data FaskesCore) (resp FaskesCore, err error)
	SelectAllFaskes() (resp []FaskesCore)
}
