package requestvaksin

type RequestvaksinCore struct {
	ID            int
	Nama          string
	NIK           int
	JenisKelamin  string
	TanggalLahir  string
	Nomor         string
	Alamat        string
	Provinsi      string
	Kota          string
	Kecamatan     string
	Kelurahan     string
	UserID        int
	VaksinID_satu int
	Status_satu   string
	Vaksin_satu   VaksinCore
	VaksinID_dua  int
	Status_dua    string
	Vaksin_dua    VaksinCore
}

type VaksinCore struct {
	ID          int
	Jenisvaksin string
	Jadwal      string
	Waktu       string
	FaskesID    int
	Faskes      FaskesCore
}

type FaskesCore struct {
	ID   int
	Nama string
}

type Bussiness interface {
	CreateRequestvaksin(data RequestvaksinCore) (resp RequestvaksinCore, err error)
	GetAllRequestvaksin() (resp []RequestvaksinCore)
	GetRequestvaksinByID(id int) (resp RequestvaksinCore, err error)
	GetRequestvaksinByUserID(id int) (resp []RequestvaksinCore, err error)
	// UpdateRequestVaksinSatuByID(data RequestvaksinCore) (resp RequestvaksinCore, err error)
}

type Data interface {
	InsertRequestvaksin(data RequestvaksinCore) (resp RequestvaksinCore, err error)
	SelectAllRequestvaksin() (resp []RequestvaksinCore)
	SelectRequestvaksinByID(id int) (resp RequestvaksinCore, err error)
	SelectRequestvaksinByUserID(id int) (resp []RequestvaksinCore, err error)
	// EditRequestVaksinSatuByID(data RequestvaksinCore) (resp RequestvaksinCore, err error)
}
