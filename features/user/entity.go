package user

type UserCore struct {
	UserID        int
	Email         string
	Namalengkap   string
	NIK           int
	Password      string
	Token         string
	RequestVaksin []RequestvaksinCore
}

type RequestvaksinCore struct {
	ID            int
	Nama          string
	NIK           int
	JenisKelamin  string
	TanggalLahir  string
	Nomor         string
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
	CreateUser(data UserCore) (resp UserCore, err error) //Register
	Login(data UserCore) (resp UserCore, err error)
	GetUserByID(id int) (resp UserCore, err error)
	UpdateUser(data UserCore) (resp UserCore, err error)
}

type Data interface {
	SelectUserEmail(data UserCore) (resp UserCore, err error)
	InsertUser(data UserCore) (resp UserCore, err error)
	Login(data UserCore) (resp UserCore, err error)
	SelectUserByID(id int) (resp UserCore, err error)
	EditUser(data UserCore) (resp UserCore, err error)
}
