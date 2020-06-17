package userMemory

const (
	UserName = "operator"
	UserPass = "operator"
)

var BasicUser User

type User struct {
	Name string
	Pass string
}

func SetUserPassword() {
	BasicUser = User{
		Name: UserName,
		Pass: UserPass,
	}
}

func UpdateUserPass(pass string) {
	BasicUser = User{
		Name: UserName,
		Pass: pass,
	}
}
