package requestvaksinasi

type RequestvaksinCore struct {
	ID              int
	Jenisvaksin     string
	Dosis_pertama   bool
	Dosis_kedua     bool
	Status_satu     string
	Status_dua      string
	Nama            string
	NIK             int
	Jam_pertama     string
	Jam_Kedua       string
	Tanggal_pertama string
	Tanggal_kedua   string
	Provinsi        string
	Kota            string
	Kelurahan       string
	Kecamatan       string
}

type Bussiness interface {
	CreateRequestvaksin(data RequestvaksinCore) (resp RequestvaksinCore, err error)
	GetAllRequestvaksin() (resp []RequestvaksinCore)
	GetRequestvaksinByID(id int) (resp RequestvaksinCore, err error)
}

type Data interface {
	InsertRequestvaksin(data RequestvaksinCore) (resp RequestvaksinCore, err error)
	SelectAllRequestvaksin() (resp []RequestvaksinCore)
	SelectRequestvaksinByID(id int) (resp RequestvaksinCore, err error)
}
