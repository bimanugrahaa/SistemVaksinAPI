package user

//buat data
type UserCore struct {
	UserID      int
	Email       string
	Namalengkap string
	NIK         int
	Password    string
	Token       string
}

type Bussiness interface {
	CreateUser(data UserCore) (resp UserCore, err error) //Register
	Login(data UserCore) (resp UserCore, err error)
}

type Data interface {
	SelectUserEmail(data UserCore) (resp UserCore, err error)
	InsertUser(data UserCore) (resp UserCore, err error)
	Login(data UserCore) (resp UserCore, err error)
}
