package admin

//buat data
type AdminCore struct {
	AdminID  int
	Username string
	Password string
	Token    string
}

type Bussiness interface {
	CreateAdmin(data AdminCore) (resp AdminCore, err error) //Register
	//Login(data AdminCore) (resp AdminCore, err error)
}

type Data interface {
	InsertAdmin(data AdminCore) (resp AdminCore, err error)
	//Login(data AdminCore) (resp AdminCore, err error)
}
